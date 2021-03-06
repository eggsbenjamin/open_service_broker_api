// +build unit

package repository_test

import (
	"database/sql"
	"testing"

	"github.com/eggsbenjamin/open_service_broker_api/db"
	"github.com/eggsbenjamin/open_service_broker_api/models"
	"github.com/eggsbenjamin/open_service_broker_api/repository"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestServiceRepository(t *testing.T) {
	t.Run("GetByID", func(t *testing.T) {
		t.Run("db error", func(t *testing.T) {
			dummyErr := errors.New("error")
			ctrl := gomock.NewController(t)
			mockDB := db.NewMockDB(ctrl)
			mockDB.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return(dummyErr)
			mockServicePlanRepo := repository.NewMockServicePlanRepository(ctrl)

			repo := repository.NewServiceRepository(mockDB, mockServicePlanRepo)
			_, err := repo.GetByServiceID("")
			require.Equal(t, dummyErr, errors.Cause(err))
			ctrl.Finish()
		})

		t.Run("not found", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockDB := db.NewMockDB(ctrl)
			mockDB.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return(sql.ErrNoRows)
			mockServicePlanRepo := repository.NewMockServicePlanRepository(ctrl)

			repo := repository.NewServiceRepository(mockDB, mockServicePlanRepo)
			_, err := repo.GetByServiceID("")
			require.Equal(t, repository.ErrNotFound, errors.Cause(err))
			ctrl.Finish()
		})
	})

	t.Run("Create", func(t *testing.T) {
		dummyErr := errors.New("error")
		ctrl := gomock.NewController(t)
		mockDB := db.NewMockDB(ctrl)
		mockDB.EXPECT().NamedQuery(gomock.Any(), gomock.Any()).Return(nil, dummyErr)
		mockServicePlanRepo := repository.NewMockServicePlanRepository(ctrl)

		serviceRepo := repository.NewServiceRepository(mockDB, mockServicePlanRepo)
		err := serviceRepo.Create(&models.DBService{})
		require.Equal(t, dummyErr, errors.Cause(err))
		ctrl.Finish()
	})
}
