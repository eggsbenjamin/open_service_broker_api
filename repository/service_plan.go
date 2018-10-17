//go:generate mockgen -package repository -source=service_plan.go -destination service_plan_mock.go

package repository

import (
	"database/sql"

	"github.com/eggsbenjamin/open_service_broker_api/db"
	"github.com/eggsbenjamin/open_service_broker_api/models"
	"github.com/pkg/errors"
)

type ServicePlanRepository interface {
	GetByServiceID(int) ([]*models.DBServicePlan, error)
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

func (s *servicePlanRepository) GetByServiceID(id int) ([]*models.DBServicePlan, error) {
	out := []*models.DBServicePlan{}
	if err := s.db.Select(
		&out,
		`
			SELECT 
				id, service_id, name
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

func (s *servicePlanRepository) Create(in *models.DBServicePlan) error {
	rows, err := s.db.NamedQuery(`
			INSERT INTO service_plan (
				service_id, name
			)
			VALUES (
				:service_id, :name 
			)
			RETURNING *
		`,
		in,
	)
	if err != nil {
		return errors.Wrap(err, "Create: error creating service plan")
	}
	defer rows.Close()

	if !rows.Next() {
		return errors.New("Create: no write result")
	}

	return rows.StructScan(in)
}