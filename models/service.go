package models

import (
	"encoding/json"

	"github.com/eggsbenjamin/open_service_broker_api/uuid"
)

type DBService struct {
	ID          int             `db:"id"`
	ServiceID   uuid.UUID       `db:"service_id"`
	Name        string          `db:"name"`
	Description string          `db:"description"`
	Tags        json.RawMessage `db:"tags"`
	Requires    json.RawMessage `db:"requires"`
	Plans       []*DBServicePlan
}

type Service struct {
	ID          uuid.UUID       `json:"service_id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Tags        json.RawMessage `json:"tags"`
	Requires    json.RawMessage `json:"requires"`
	Plans       []*ServicePlan  `json:"plans"`
}
