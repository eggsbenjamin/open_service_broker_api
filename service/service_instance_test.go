package service_test

import (
	"testing"

	models "github.com/eggsbenjamin/open_service_broker_api/models"
	"github.com/eggsbenjamin/open_service_broker_api/repository"
	"github.com/eggsbenjamin/open_service_broker_api/service"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestServiceInstanceService(t *testing.T) {
	t.Run("serviceRepo GetByServiceID() error", func(t *testing.T) {
		t.Run("unexpected", func(t *testing.T) {
			dummyErr := errors.New("error")
			ctrl := gomock.NewController(t)
			mockServiceRepo := repository.NewMockServiceRepository(ctrl)
			mockServicePlanRepo := repository.NewMockServicePlanRepository(ctrl)
			mockServiceInstanceRepo := repository.NewMockServiceInstanceRepository(ctrl)
			mockServiceRepo.EXPECT().GetByServiceID(gomock.Any()).Return(nil, dummyErr)

			serviceInstanceSrv := service.NewServiceInstanceService(
				mockServiceRepo,
				mockServicePlanRepo,
				mockServiceInstanceRepo,
			)
			err := serviceInstanceSrv.Create(&models.ServiceInstance{})
			require.Equal(t, dummyErr, errors.Cause(err))
			ctrl.Finish()
		})

		t.Run("ErrNotFound", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockServiceRepo := repository.NewMockServiceRepository(ctrl)
			mockServicePlanRepo := repository.NewMockServicePlanRepository(ctrl)
			mockServiceInstanceRepo := repository.NewMockServiceInstanceRepository(ctrl)
			mockServiceRepo.EXPECT().GetByServiceID(gomock.Any()).Return(nil, repository.ErrNotFound)

			serviceInstanceSrv := service.NewServiceInstanceService(
				mockServiceRepo,
				mockServicePlanRepo,
				mockServiceInstanceRepo,
			)
			err := serviceInstanceSrv.Create(&models.ServiceInstance{})
			require.Equal(t, service.ErrServiceNotFound, err)
			ctrl.Finish()
		})
	})

	t.Run("servicePlanRepo GetByPlanID() error", func(t *testing.T) {
		t.Run("unexpected", func(t *testing.T) {
			dummyErr := errors.New("error")
			ctrl := gomock.NewController(t)
			mockServiceRepo := repository.NewMockServiceRepository(ctrl)
			mockServicePlanRepo := repository.NewMockServicePlanRepository(ctrl)
			mockServiceInstanceRepo := repository.NewMockServiceInstanceRepository(ctrl)
			mockServiceRepo.EXPECT().GetByServiceID(gomock.Any()).Return(nil, nil)
			mockServicePlanRepo.EXPECT().GetByPlanID(gomock.Any()).Return(nil, dummyErr)

			serviceInstanceSrv := service.NewServiceInstanceService(
				mockServiceRepo,
				mockServicePlanRepo,
				mockServiceInstanceRepo,
			)
			err := serviceInstanceSrv.Create(&models.ServiceInstance{})
			require.Equal(t, dummyErr, errors.Cause(err))
			ctrl.Finish()
		})

		t.Run("ErrNotFound", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockServiceRepo := repository.NewMockServiceRepository(ctrl)
			mockServicePlanRepo := repository.NewMockServicePlanRepository(ctrl)
			mockServiceInstanceRepo := repository.NewMockServiceInstanceRepository(ctrl)
			mockServiceRepo.EXPECT().GetByServiceID(gomock.Any()).Return(nil, nil)
			mockServicePlanRepo.EXPECT().GetByPlanID(gomock.Any()).Return(nil, repository.ErrNotFound)

			serviceInstanceSrv := service.NewServiceInstanceService(
				mockServiceRepo,
				mockServicePlanRepo,
				mockServiceInstanceRepo,
			)
			err := serviceInstanceSrv.Create(&models.ServiceInstance{})
			require.Equal(t, service.ErrPlanNotFound, err)
			ctrl.Finish()
		})
	})

	t.Run("serviceInstanceRepo Create() error", func(t *testing.T) {
		dummyErr := errors.New("error")
		ctrl := gomock.NewController(t)
		mockServiceRepo := repository.NewMockServiceRepository(ctrl)
		mockServicePlanRepo := repository.NewMockServicePlanRepository(ctrl)
		mockServiceInstanceRepo := repository.NewMockServiceInstanceRepository(ctrl)
		mockServiceRepo.EXPECT().GetByServiceID(gomock.Any()).Return(nil, nil)
		mockServicePlanRepo.EXPECT().GetByPlanID(gomock.Any()).Return(&models.DBServicePlan{}, nil)
		mockServiceInstanceRepo.EXPECT().Create(gomock.Any()).Return(dummyErr)

		serviceInstanceSrv := service.NewServiceInstanceService(
			mockServiceRepo,
			mockServicePlanRepo,
			mockServiceInstanceRepo,
		)
		err := serviceInstanceSrv.Create(&models.ServiceInstance{})
		require.Equal(t, dummyErr, errors.Cause(err))
		ctrl.Finish()
	})
}
