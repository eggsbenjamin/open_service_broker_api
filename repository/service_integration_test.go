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
		db, err := db.NewConnection("localhost", "1234", "postgres", "postgres", "service_catalog")
		require.NoError(t, err)
		require.NoError(t, testutils.TeardownDB(db))
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

		servicePlans := []*models.DBServicePlan{
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

	t.Run("GetAll", func(t *testing.T) {
		db, err := db.NewConnection("localhost", "1234", "postgres", "postgres", "service_catalog")
		require.NoError(t, err)
		require.NoError(t, testutils.TeardownDB(db))
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

		servicePlans := []*models.DBServicePlan{
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

		for _, servicePlan := range servicePlans {
			require.NoError(t, servicePlanRepo.Create(servicePlan))
		}

		service1.Plans = append(service1.Plans, servicePlans[0])
		service1.Plans = append(service1.Plans, servicePlans[1])
		service2.Plans = append(service2.Plans, servicePlans[2])

		expectedServices := []*models.DBService{
			service1,
			service2,
		}

		services, err := serviceRepo.GetAll()
		require.NoError(t, err)
		require.Equal(t, len(expectedServices), len(services))

		for id := range services {
			require.Equal(t, expectedServices[id], services[id])
		}
	})
}
