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

func TestServicePlanRepositoryIntegration(t *testing.T) {
	t.Run("GetByServiceID", func(t *testing.T) {
		db, err := db.NewConnection("localhost", "32768", "postgres", "postgres", "service_catalog")
		require.NoError(t, testutils.TeardownDB(db))
		require.NoError(t, err)
		defer testutils.TeardownDB(db)

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

		input := []*models.DBServicePlan{
			{
				ServiceID: service.ID,
				Name:      "test1",
			},
			{
				ServiceID: service.ID,
				Name:      "test2",
			},
		}

		for _, in := range input {
			require.NoError(t, servicePlanRepo.Create(in))
		}

		servicePlans, err := servicePlanRepo.GetByServiceID(service.ID)
		require.NoError(t, err)
		require.Equal(t, len(input), len(servicePlans))

		for i := range servicePlans {
			require.Equal(t, input[i], servicePlans[i])
		}
	})
}
