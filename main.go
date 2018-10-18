package main

import (
	"log"
	"net/http"
	"os"

	"github.com/eggsbenjamin/open_service_broker_api/db"
	"github.com/eggsbenjamin/open_service_broker_api/handlers"
	"github.com/eggsbenjamin/open_service_broker_api/repository"
	"github.com/eggsbenjamin/open_service_broker_api/service"
	"github.com/go-chi/chi"
)

func main() {
	db, err := db.NewConnection(
		mustGetEnv("DB_HOST"),
		mustGetEnv("DB_PORT"),
		mustGetEnv("DB_USER"),
		mustGetEnv("DB_PWD"),
		mustGetEnv("DB_NAME"),
	)
	if err != nil {
		log.Fatalf("error setting up db: %q", err)
	}

	// repos
	servicePlanRepo := repository.NewServicePlanRepository(db)
	serviceRepo := repository.NewServiceRepository(db, servicePlanRepo)

	// services
	catalogService := service.NewCatalogService(serviceRepo)

	// handlers
	catalogHandlers := handlers.NewCatalogHandlers(catalogService)

	srv := chi.NewMux()
	srv.Route("/v1", func(r chi.Router) {
		r.Get("/catalog", catalogHandlers.GetCatalog)
	})

	log.Fatal(http.ListenAndServe(":8080", srv))
}

func mustGetEnv(k string) string {
	v, ok := os.LookupEnv(k)
	if !ok {
		log.Fatalf("env var %s not set", k)
	}

	return v
}
