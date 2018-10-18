package models

import (
	"encoding/json"

	"github.com/eggsbenjamin/open_service_broker_api/uuid"
)

type DBServiceInstance struct {
	ID     int             `db:"id"`
	PlanID int             `db:"plan_id"`
	Ctx    json.RawMessage `db:"context"`
	Params json.RawMessage `db:"parameters"`
}

type ServiceInstance struct {
	ID        uuid.UUID       `json:"id"`
	ServiceID uuid.UUID       `json:"service_id"`
	PlanID    uuid.UUID       `json:"plan_id"`
	Ctx       json.RawMessage `json:"context"`
	Params    json.RawMessage `json:"parameters"`
}
