//go:generate mockgen -package service -source=service_instance.go -destination service_instance_mock.go

package service

import (
	models "github.com/eggsbenjamin/open_service_broker_api/models"
	"github.com/eggsbenjamin/open_service_broker_api/repository"
	"github.com/pkg/errors"
)

type ServiceInstanceService interface {
	Create(*models.ServiceInstance) error
}

type serviceInstanceService struct {
	serviceRepo         repository.ServiceRepository
	servicePlanRepo     repository.ServicePlanRepository
	serviceInstanceRepo repository.ServiceInstanceRepository
}

func NewServiceInstanceService(
	serviceRepo repository.ServiceRepository,
	servicePlanRepo repository.ServicePlanRepository,
	serviceInstanceRepo repository.ServiceInstanceRepository,
) ServiceInstanceService {
	return &serviceInstanceService{
		serviceRepo:         serviceRepo,
		servicePlanRepo:     servicePlanRepo,
		serviceInstanceRepo: serviceInstanceRepo,
	}
}

func (s *serviceInstanceService) Create(serviceInstance *models.ServiceInstance) error {
	_, err := s.serviceRepo.GetByServiceID(serviceInstance.ServiceID)
	if err != nil {
		if err == repository.ErrNotFound {
			return ErrServiceNotFound
		}
		return errors.Wrap(err, "Create: error getting service")
	}

	servicePlan, err := s.servicePlanRepo.GetByPlanID(serviceInstance.PlanID)
	if err != nil {
		if err == repository.ErrNotFound {
			return ErrPlanNotFound
		}
		return errors.Wrap(err, "Create: error getting service plan")
	}

	return s.serviceInstanceRepo.Create(
		mapDBServiceInstance(serviceInstance, servicePlan.ID),
	)
}

func mapDBServiceInstance(serviceInstance *models.ServiceInstance, planID int) *models.DBServiceInstance {
	return &models.DBServiceInstance{
		PlanID: planID,
		Ctx:    serviceInstance.Ctx,
		Params: serviceInstance.Params,
	}
}
