package models

import "github.com/eggsbenjamin/open_service_broker_api/uuid"

type DBServicePlan struct {
	ID        int       `db:"id"`
	ServiceID int       `db:"service_id"`
	PlanID    uuid.UUID `db:"plan_id"`
	Name      string    `db:"name"`
}

type ServicePlan struct {
	ServiceID uuid.UUID `json:"service_id"`
	Name      string    `json:"name"`
}
