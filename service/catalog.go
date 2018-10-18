package service

import (
	"github.com/eggsbenjamin/open_service_broker_api/models"
	"github.com/eggsbenjamin/open_service_broker_api/repository"
	"github.com/pkg/errors"
)

type CatalogService interface {
	GetCatalog() (*models.Catalog, error)
}

type catalogService struct {
	serviceRepo repository.ServiceRepository
}

func NewCatalogService(serviceRepo repository.ServiceRepository) CatalogService {
	return &catalogService{
		serviceRepo: serviceRepo,
	}
}

func (c *catalogService) GetCatalog() (*models.Catalog, error) {
	services, err := c.serviceRepo.GetAll()
	if err != nil {
		return nil, errors.Wrap(err, "GetCatalog: error getting services")
	}

	catalog := &models.Catalog{}
	for _, service := range services {
		catalog.Services = append(catalog.Services, mapService(service))
	}

	return catalog, nil
}

func mapService(dbService *models.DBService) *models.Service {
	return &models.Service{
		ID:          dbService.ServiceID,
		Name:        dbService.Name,
		Description: dbService.Description,
		Tags:        dbService.Tags,
		Requires:    dbService.Requires,
	}
}
