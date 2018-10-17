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

func TestServiceRepositoryIntegration(t *testing.T) {
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

		servicePlans := []*models.DBServicePlan{
			{
				ServiceID: service.ID,
				Name:      "test1",
			},
			{
				ServiceID: service.ID,
				Name:      "test2",
			},
		}

		for _, servicePlan := range servicePlans {
			require.NoError(t, servicePlanRepo.Create(servicePlan))
		}

		result, err := serviceRepo.GetByServiceID(id)
		require.NoError(t, err)
		require.NotZero(t, result.ID)
		require.Equal(t, service.Name, result.Name)
		require.Equal(t, service.Description, result.Description)
		require.JSONEq(t, string(service.Tags), string(result.Tags))
		require.JSONEq(t, string(service.Requires), string(result.Requires))
		require.Equal(t, len(servicePlans), len(result.Plans))
	})
}
