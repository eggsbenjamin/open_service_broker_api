//go:generate mockgen -package repository -source=service.go -destination service_mock.go

package repository

import (
	"database/sql"

	"github.com/eggsbenjamin/open_service_broker_api/db"
	"github.com/eggsbenjamin/open_service_broker_api/models"
	"github.com/eggsbenjamin/open_service_broker_api/uuid"
	"github.com/pkg/errors"
)

type ServiceRepository interface {
	GetAll() ([]*models.DBService, error)
	GetByServiceID(uuid.UUID) (*models.DBService, error)
	Create(*models.DBService) error
}

type serviceRepository struct {
	db              db.DB
	servicePlanRepo ServicePlanRepository
}

func NewServiceRepository(db db.DB, servicePlanRepo ServicePlanRepository) ServiceRepository {
	return &serviceRepository{
		db:              db,
		servicePlanRepo: servicePlanRepo,
	}
}

func (s *serviceRepository) GetAll() ([]*models.DBService, error) {
	out := []*models.DBService{}
	if err := s.db.Select(
		&out,
		`
			SELECT 
				id, service_id, name, description, tags, requires
			FROM service
		`,
	); err != nil {
		return nil, err
	}

	if len(out) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "GetAll: no results")
	}

	serviceIDs := []int{}
	for _, service := range out {
		serviceIDs = append(serviceIDs, service.ID)
	}

	servicePlans, err := s.servicePlanRepo.GetByServiceIDs(serviceIDs...)
	if err != nil {
		return nil, err
	}

	for _, service := range out {
		service.Plans = append(service.Plans, servicePlans[service.ID]...)
	}

	return out, nil
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

	servicePlans, err := s.servicePlanRepo.GetByServiceID(out.ID)
	if err != nil {
		return nil, err
	}

	out.Plans = servicePlans
	return out, nil
}

func (s *serviceRepository) Create(in *models.DBService) error {
	rows, err := s.db.NamedQuery(`
		INSERT INTO service (
			service_id, name, description, tags, requires
		)
		VALUES (
			:service_id, :name, :description, :tags, :requires
		)
		RETURNING *;
	`,
		in,
	)
	if err != nil {
		return errors.Wrap(err, "Create: error creating service")
	}
	defer rows.Close()

	if !rows.Next() {
		return errors.New("Create: no write result")
	}

	return rows.StructScan(in)
}
