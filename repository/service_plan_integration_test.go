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

		id2, err := uuid.New()
		require.NoError(t, err)

		id3, err := uuid.New()
		require.NoError(t, err)

		input := []*models.DBServicePlan{
			{
				PlanID:    id2,
				ServiceID: service.ID,
				Name:      "test1",
			},
			{
				PlanID:    id3,
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

	t.Run("GetByServiceIDs", func(t *testing.T) {
		db, err := db.NewConnection("localhost", "32768", "postgres", "postgres", "service_catalog")
		require.NoError(t, testutils.TeardownDB(db))
		require.NoError(t, err)
		defer testutils.TeardownDB(db)

		servicePlanRepo := repository.NewServicePlanRepository(db)
		serviceRepo := repository.NewServiceRepository(db, servicePlanRepo)

		id1, err := uuid.New()
		require.NoError(t, err)

		id2, err := uuid.New()
		require.NoError(t, err)

		service1 := &models.DBService{
			ServiceID:   id1,
			Name:        "test1",
			Description: "test service1",
			Tags:        []byte(`["one", "two", "three"]`),
			Requires:    []byte(`{"one": "two"}`),
		}
		require.NoError(t, serviceRepo.Create(service1))

		service2 := &models.DBService{
			ServiceID:   id2,
			Name:        "test2",
			Description: "test service2",
			Tags:        []byte(`["one", "two", "three"]`),
			Requires:    []byte(`{"one": "two"}`),
		}
		require.NoError(t, serviceRepo.Create(service2))

		id3, err := uuid.New()
		require.NoError(t, err)

		id4, err := uuid.New()
		require.NoError(t, err)

		id5, err := uuid.New()
		require.NoError(t, err)

		input := []*models.DBServicePlan{
			{
				PlanID:    id3,
				ServiceID: service1.ID,
				Name:      "test1",
			},
			{
				PlanID:    id4,
				ServiceID: service1.ID,
				Name:      "test2",
			},
			{
				PlanID:    id5,
				ServiceID: service2.ID,
				Name:      "test3",
			},
		}

		expectedServicePlans := map[int][]*models.DBServicePlan{
			service1.ID: {
				input[0],
				input[1],
			},
			service2.ID: {
				input[2],
			},
		}

		for _, in := range input {
			require.NoError(t, servicePlanRepo.Create(in))
		}

		servicePlans, err := servicePlanRepo.GetByServiceIDs(service1.ID, service2.ID)
		require.NoError(t, err)
		require.Equal(t, len(expectedServicePlans), len(servicePlans))

		for k := range servicePlans {
			require.ElementsMatch(t, expectedServicePlans[k], servicePlans[k])
		}
	})
}
