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

func TestServicePlanRepository(t *testing.T) {
	t.Run("GetByServiceID", func(t *testing.T) {
		t.Run("db error", func(t *testing.T) {
			dummyErr := errors.New("error")
			ctrl := gomock.NewController(t)
			mockDB := db.NewMockDB(ctrl)
			mockDB.EXPECT().Select(gomock.Any(), gomock.Any(), gomock.Any()).Return(dummyErr)

			repo := repository.NewServicePlanRepository(mockDB)
			_, err := repo.GetByServiceID(1)
			require.Equal(t, dummyErr, errors.Cause(err))
			ctrl.Finish()
		})

		t.Run("not found", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockDB := db.NewMockDB(ctrl)
			mockDB.EXPECT().Select(gomock.Any(), gomock.Any(), gomock.Any()).Return(sql.ErrNoRows)

			repo := repository.NewServicePlanRepository(mockDB)
			_, err := repo.GetByServiceID(1)
			require.Equal(t, repository.ErrNotFound, errors.Cause(err))
			ctrl.Finish()
		})
	})

	t.Run("Create", func(t *testing.T) {
		dummyErr := errors.New("error")
		ctrl := gomock.NewController(t)
		mockDB := db.NewMockDB(ctrl)
		mockDB.EXPECT().NamedQuery(gomock.Any(), gomock.Any()).Return(nil, dummyErr)

		repo := repository.NewServicePlanRepository(mockDB)
		err := repo.Create(&models.DBServicePlan{})
		require.Equal(t, dummyErr, errors.Cause(err))
		ctrl.Finish()
	})
}
