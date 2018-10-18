package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eggsbenjamin/open_service_broker_api/handlers"
	"github.com/eggsbenjamin/open_service_broker_api/service"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestCatalogHandlers(t *testing.T) {
	t.Run("catalogService.GetAll() error", func(t *testing.T) {
		dummyErr := errors.New("error")
		ctrl := gomock.NewController(t)
		mockCatalogService := service.NewMockCatalogService(ctrl)
		mockCatalogService.EXPECT().GetCatalog().Return(nil, dummyErr)

		catalogHandlers := handlers.NewCatalogHandlers(mockCatalogService)
		req := httptest.NewRequest(http.MethodGet, "/v1/catalog", nil)
		rec := httptest.NewRecorder()

		catalogHandlers.GetCatalog(rec, req)
		res := rec.Result()
		require.Equal(t, res.StatusCode, http.StatusInternalServerError)
		ctrl.Finish()
	})
}
