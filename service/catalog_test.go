// +build unit

package service_test

import (
	"testing"

	"github.com/eggsbenjamin/open_service_broker_api/models"
	"github.com/eggsbenjamin/open_service_broker_api/repository"
	"github.com/eggsbenjamin/open_service_broker_api/service"
	"github.com/eggsbenjamin/open_service_broker_api/uuid"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestCatalogService(t *testing.T) {
	t.Run("GetCatalog", func(t *testing.T) {
		t.Run("repo.GetAll error", func(t *testing.T) {
			dummyErr := errors.New("error")
			ctrl := gomock.NewController(t)
			mockServiceRepo := repository.NewMockServiceRepository(ctrl)
			mockServiceRepo.EXPECT().GetAll().Return(nil, dummyErr)

			catalogSrv := service.NewCatalogService(mockServiceRepo)
			_, err := catalogSrv.GetCatalog()
			require.Equal(t, dummyErr, errors.Cause(err))
			ctrl.Finish()
		})

		t.Run("success", func(t *testing.T) {
			id, err := uuid.New()
			require.NoError(t, err)

			ctrl := gomock.NewController(t)
			mockServiceRepo := repository.NewMockServiceRepository(ctrl)
			mockServiceRepo.EXPECT().GetAll().Return([]*models.DBService{
				{

					ServiceID:   id,
					Name:        "test2",
					Description: "test service2",
					Tags:        []byte(`["one", "two", "three"]`),
					Requires:    []byte(`{"one": "two"}`),
				},
			}, nil)

			expectedCatalog := &models.Catalog{
				Services: []*models.Service{
					{
						ID:          id,
						Name:        "test2",
						Description: "test service2",
						Tags:        []byte(`["one", "two", "three"]`),
						Requires:    []byte(`{"one": "two"}`),
					},
				},
			}

			catalogSrv := service.NewCatalogService(mockServiceRepo)
			catalog, err := catalogSrv.GetCatalog()
			require.NoError(t, err)
			require.Equal(t, expectedCatalog, catalog)
			ctrl.Finish()
		})
	})
}
