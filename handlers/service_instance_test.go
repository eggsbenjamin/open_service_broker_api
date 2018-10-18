package handlers_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eggsbenjamin/open_service_broker_api/handlers"
	"github.com/eggsbenjamin/open_service_broker_api/service"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestServiceInstanceHandlers(t *testing.T) {
	t.Run("serviceInstanceService.Create() error", func(t *testing.T) {
		t.Run("unexpected", func(t *testing.T) {
			dummyErr := errors.New("error")
			ctrl := gomock.NewController(t)
			mockServiceInstanceService := service.NewMockServiceInstanceService(ctrl)
			mockServiceInstanceService.EXPECT().Create(gomock.Any()).Return(dummyErr)

			serviceInstanceHandlers := handlers.NewServiceInstanceHandlers(mockServiceInstanceService)
			body := ioutil.NopCloser(bytes.NewBufferString(`{}`))
			req := httptest.NewRequest(http.MethodPut, "/v1/service_instances/12345", body)
			rec := httptest.NewRecorder()

			serviceInstanceHandlers.CreateServiceInstance(rec, req)
			res := rec.Result()
			require.Equal(t, res.StatusCode, http.StatusInternalServerError)
			ctrl.Finish()
		})

		t.Run("unknown service", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockServiceInstanceService := service.NewMockServiceInstanceService(ctrl)
			mockServiceInstanceService.EXPECT().Create(gomock.Any()).Return(service.ErrServiceNotFound)

			serviceInstanceHandlers := handlers.NewServiceInstanceHandlers(mockServiceInstanceService)
			body := ioutil.NopCloser(bytes.NewBufferString(`{}`))
			req := httptest.NewRequest(http.MethodPut, "/v1/service_instances/12345", body)
			rec := httptest.NewRecorder()

			serviceInstanceHandlers.CreateServiceInstance(rec, req)
			res := rec.Result()
			require.Equal(t, res.StatusCode, http.StatusBadRequest)
			ctrl.Finish()
		})

		t.Run("unknown plan", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockServiceInstanceService := service.NewMockServiceInstanceService(ctrl)
			mockServiceInstanceService.EXPECT().Create(gomock.Any()).Return(service.ErrPlanNotFound)

			serviceInstanceHandlers := handlers.NewServiceInstanceHandlers(mockServiceInstanceService)
			body := ioutil.NopCloser(bytes.NewBufferString(`{}`))
			req := httptest.NewRequest(http.MethodPut, "/v1/service_instances/12345", body)
			rec := httptest.NewRecorder()

			serviceInstanceHandlers.CreateServiceInstance(rec, req)
			res := rec.Result()
			require.Equal(t, res.StatusCode, http.StatusBadRequest)
			ctrl.Finish()
		})
	})
}
