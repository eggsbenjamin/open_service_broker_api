//go:generate mockgen -package repository -source=service_plan.go -destination service_plan_mock.go

package repository

import (
	"database/sql"

	"github.com/eggsbenjamin/open_service_broker_api/db"
	"github.com/eggsbenjamin/open_service_broker_api/models"
	uuid "github.com/eggsbenjamin/open_service_broker_api/uuid"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type ServicePlanRepository interface {
	GetByPlanID(uuid.UUID) (*models.DBServicePlan, error)
	GetByServiceID(int) ([]*models.DBServicePlan, error)
	GetByServiceIDs(...int) (map[int][]*models.DBServicePlan, error)
	Create(*models.DBServicePlan) error
}

type servicePlanRepository struct {
	db db.DB
}

func NewServicePlanRepository(db db.DB) ServicePlanRepository {
	return &servicePlanRepository{
		db: db,
	}
}

func (s *servicePlanRepository) GetByPlanID(id uuid.UUID) (*models.DBServicePlan, error) {
	out := &models.DBServicePlan{}
	if err := s.db.Get(
		out,
		`
			SELECT 
				id, plan_id, service_id, name
			FROM service_plan
			WHERE plan_id = $1
		`,
		id,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrapf(ErrNotFound, "GetByID: no results for id: %s", id)
		}
		return nil, err
	}

	return out, nil
}

func (s *servicePlanRepository) GetByServiceID(id int) ([]*models.DBServicePlan, error) {
	out := []*models.DBServicePlan{}
	if err := s.db.Select(
		&out,
		`
			SELECT 
				id, plan_id, service_id, name
			FROM service_plan
			WHERE service_id = $1
		`,
		id,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrapf(ErrNotFound, "GetByID: no results for id: %s", id)
		}
		return nil, err
	}

	return out, nil
}

func (s *servicePlanRepository) GetByServiceIDs(ids ...int) (map[int][]*models.DBServicePlan, error) {
	out := map[int][]*models.DBServicePlan{}
	rows, err := s.db.Query(
		`
			SELECT 
				id, plan_id, service_id, name
			FROM service_plan
			WHERE service_id = ANY($1)
		`,
		pq.Array(ids),
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrapf(ErrNotFound, "GetByIDs: no results for id: %q", ids)
		}
		return nil, err
	}
	defer rows.Close() // nolint: errcheck

	for rows.Next() {
		servicePlan := &models.DBServicePlan{}
		if err := rows.StructScan(servicePlan); err != nil {
			return nil, errors.Wrap(err, "GetByIDs: error scanning service plan values")
		}
		out[servicePlan.ServiceID] = append(out[servicePlan.ServiceID], servicePlan)
	}

	return out, nil
}

func (s *servicePlanRepository) Create(in *models.DBServicePlan) error {
	rows, err := s.db.NamedQuery(`
			INSERT INTO service_plan (
				plan_id, service_id, name
			)
			VALUES (
				:plan_id, :service_id, :name 
			)
			RETURNING *
		`,
		in,
	)
	if err != nil {
		return errors.Wrap(err, "Create: error creating service plan")
	}
	defer rows.Close() // nolint: errcheck

	if !rows.Next() {
		return errors.New("Create: no write result")
	}

	return rows.StructScan(in)
}
