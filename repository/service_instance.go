//go:generate mockgen -package repository -source=service_instance.go -destination service_instance_mock.go

package repository

import (
	"database/sql"

	"github.com/eggsbenjamin/open_service_broker_api/db"
	"github.com/eggsbenjamin/open_service_broker_api/models"
	"github.com/pkg/errors"
)

type ServiceInstanceRepository interface {
	GetByID(int) (*models.DBServiceInstance, error)
	Create(*models.DBServiceInstance) error
}

type serviceInstanceRepository struct {
	db db.DB
}

func NewServiceInstanceRepository(db db.DB) ServiceInstanceRepository {
	return &serviceInstanceRepository{
		db: db,
	}
}

func (s *serviceInstanceRepository) GetByID(id int) (*models.DBServiceInstance, error) {
	out := &models.DBServiceInstance{}
	if err := s.db.Get(
		out,
		`
			SELECT 
				id, plan_id, context, parameters
			FROM service_instance
			WHERE id = $1
		`,
		id,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrapf(ErrNotFound, "GetByID: no results for id: %d", id)
		}
		return nil, err
	}

	return out, nil
}

func (s *serviceInstanceRepository) Create(in *models.DBServiceInstance) error {
	rows, err := s.db.NamedQuery(`
			INSERT INTO service_instance (
				plan_id, context, parameters
			)
			VALUES (
				:plan_id, :context, :parameters
			)
			RETURNING *;
		`,
		in,
	)
	if err != nil {
		return errors.Wrap(err, "Create: error creating service instance")
	}
	defer rows.Close() // nolint: errcheck

	if !rows.Next() {
		return errors.New("Create: no write result")
	}

	return rows.StructScan(in)
}
