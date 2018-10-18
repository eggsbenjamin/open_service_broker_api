package handlers

import (
	"net/http"

	"github.com/eggsbenjamin/open_service_broker_api/service"
)

type CatalogHandlers struct {
	catalogService service.CatalogService
}

func (c *CatalogHandlers) GetCatalog(w http.ResponseWriter, r *http.Request) {

}
