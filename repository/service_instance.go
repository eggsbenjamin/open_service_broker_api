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
	var out *models.DBServiceInstance
	if err := s.db.Get(&out, ``, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrapf(ErrNotFound, "GetByID: no results for id: %d", id)
		}
		return nil, err
	}

	return out, nil
}

func (s *serviceInstanceRepository) Create(in *models.DBServiceInstance) error {
	if _, err := s.db.NamedQuery("", &in); err != nil {
		return errors.Wrapf(err, "Create: error creating service instance")
	}

	return nil
}
