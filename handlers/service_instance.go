package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/eggsbenjamin/open_service_broker_api/models"
	"github.com/eggsbenjamin/open_service_broker_api/service"
	"github.com/eggsbenjamin/open_service_broker_api/uuid"
)

type serviceInstanceHandlers struct {
	serviceInstanceService service.ServiceInstanceService
}

func NewServiceInstanceHandlers(serviceInstanceService service.ServiceInstanceService) *serviceInstanceHandlers {
	return &serviceInstanceHandlers{
		serviceInstanceService: serviceInstanceService,
	}
}

func (c *serviceInstanceHandlers) CreateServiceInstance(w http.ResponseWriter, r *http.Request) {
	serviceInstance := &models.ServiceInstance{}
	if err := json.NewDecoder(r.Body).Decode(serviceInstance); err != nil {
		log.Printf("error decoding service instance request data: %q", err)
		sendError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	defer r.Body.Close()

	uri := strings.Split(r.URL.Path, "/")
	id := uri[len(uri)-1]
	serviceInstance.ID = uuid.UUID(id)

	err := c.serviceInstanceService.Create(serviceInstance)
	if err != nil {
		log.Printf("error creating service instance: %q", err)
		switch err {
		case service.ErrServiceNotFound:
			sendError(w, http.StatusBadRequest, "unknown service")
		case service.ErrPlanNotFound:
			sendError(w, http.StatusBadRequest, "unknown plan")
		default:
			sendError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}
}
