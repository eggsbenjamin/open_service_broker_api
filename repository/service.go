package repository

import (
	"database/sql"

	"github.com/eggsbenjamin/open_service_broker_api/db"
	"github.com/eggsbenjamin/open_service_broker_api/models"
	"github.com/eggsbenjamin/open_service_broker_api/uuid"
	"github.com/pkg/errors"
)

type ServiceRepository interface {
	GetByServiceID(uuid.UUID) (*models.DBService, error)
	Create(*models.DBService) error
}

type serviceRepository struct {
	db db.DB
}

func NewServiceRepository(db db.DB) ServiceRepository {
	return &serviceRepository{
		db: db,
	}
}

func (s *serviceRepository) GetByServiceID(id uuid.UUID) (*models.DBService, error) {
	out := &models.DBService{}
	if err := s.db.Get(
		out,
		`
			SELECT 
				id, service_id, name, description, tags, requires
			FROM service
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

func (s *serviceRepository) Create(in *models.DBService) error {
	if _, err := s.db.NamedQuery(`
		INSERT INTO service (
			service_id, name, description, tags, requires
		)
		VALUES (
			:service_id, :name, :description, :tags, :requires
		)
	`,
		map[string]interface{}{
			"service_id":  in.ServiceID,
			"name":        in.Name,
			"description": in.Description,
			"tags":        in.Tags,
			"requires":    in.Requires,
		}); err != nil {
		return errors.Wrap(err, "Create: error creating service")
	}

	return nil
}
