package service

import (
	models "github.com/eggsbenjamin/open_service_broker_api/models"
	"github.com/eggsbenjamin/open_service_broker_api/repository"
)

type ServiceInstanceService interface {
	Create(*models.ServiceInstance) error
}

type serviceInstanceService struct {
	serviceInstanceRepo repository.ServiceInstanceRepository
}

func NewServiceInstanceService(serviceInstanceRepo repository.ServiceInstanceRepository) ServiceInstanceService {
	return &serviceInstanceService{
		serviceInstanceRepo: serviceInstanceRepo,
	}
}

func (s *serviceInstanceService) Create(*models.ServiceInstance) error {
	return nil
}
