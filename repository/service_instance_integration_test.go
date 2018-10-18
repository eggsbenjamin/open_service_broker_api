// +build integration

package repository_test

import (
	"testing"

	"github.com/eggsbenjamin/open_service_broker_api/db"
	"github.com/eggsbenjamin/open_service_broker_api/models"
	"github.com/eggsbenjamin/open_service_broker_api/repository"
	"github.com/eggsbenjamin/open_service_broker_api/testutils"
	"github.com/eggsbenjamin/open_service_broker_api/uuid"
	"github.com/stretchr/testify/require"
)

func TestServiceInstanceRepositoryIntegration(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		db, err := db.NewConnection("localhost", "32768", "postgres", "postgres", "service_catalog")
		require.NoError(t, testutils.TeardownDB(db))
		require.NoError(t, err)
		defer testutils.TeardownDB(db)

		serviceInstanceRepo := repository.NewServiceInstanceRepository(db)
		servicePlanRepo := repository.NewServicePlanRepository(db)
		serviceRepo := repository.NewServiceRepository(db, servicePlanRepo)
		id, err := uuid.New()
		require.NoError(t, err)

		service := &models.DBService{
			ServiceID:   id,
			Name:        "test",
			Description: "test service",
			Tags:        []byte(`["one", "two", "three"]`),
			Requires:    []byte(`{"one": "two"}`),
		}
		require.NoError(t, serviceRepo.Create(service))

		servicePlan := &models.DBServicePlan{
			PlanID:    id,
			ServiceID: service.ID,
			Name:      "test1",
		}
		require.NoError(t, servicePlanRepo.Create(servicePlan))

		serviceInstance := &models.DBServiceInstance{
			PlanID: servicePlan.ID,
			Ctx:    []byte(`{"context":"test"}`),
			Params: []byte(`{"parameters":"test"}`),
		}

		err = serviceInstanceRepo.Create(serviceInstance)
		require.NoError(t, err)

		result, err := serviceInstanceRepo.GetByID(serviceInstance.ID)
		require.NoError(t, err)
		require.Equal(t, serviceInstance, result)
	})
}
