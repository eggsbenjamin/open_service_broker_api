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
	t.Run("Create", func(t *testing.T) {
		db, err := db.NewConnection("localhost", "32768", "postgres", "postgres", "service_catalog")
		require.NoError(t, testutils.TeardownDB(db))
		require.NoError(t, err)
		defer testutils.TeardownDB(db)

		repo := repository.NewServiceRepository(db)
		id, err := uuid.New()
		require.NoError(t, err)

		input := &models.DBService{
			ServiceID:   id,
			Name:        "test",
			Description: "test service",
			Tags:        []byte(`["one", "two", "three"]`),
			Requires:    []byte(`{"one": "two"}`),
		}
		err = repo.Create(input)
		require.NoError(t, err)

		service, err := repo.GetByServiceID(id)
		require.NoError(t, err)
		require.NotZero(t, service.ID)
		require.Equal(t, input.Name, service.Name)
		require.Equal(t, input.Description, service.Description)
		require.JSONEq(t, string(input.Tags), string(service.Tags))
		require.JSONEq(t, string(input.Requires), string(service.Requires))
	})
}
