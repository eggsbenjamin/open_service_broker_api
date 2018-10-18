package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/eggsbenjamin/open_service_broker_api/service"
)

type catalogHandlers struct {
	catalogService service.CatalogService
}

func NewCatalogHandlers(catalogService service.CatalogService) *catalogHandlers {
	return &catalogHandlers{
		catalogService: catalogService,
	}
}

func (c *catalogHandlers) GetCatalog(w http.ResponseWriter, r *http.Request) {
	catalog, err := c.catalogService.GetCatalog()
	if err != nil {
		log.Printf("error getting catalog data: %q", err)
		sendError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	if err := json.NewEncoder(w).Encode(catalog); err != nil {
		log.Printf("error encoding catalog json: %q", err)
		sendError(w, http.StatusInternalServerError, "internal server error")
		return
	}
}
