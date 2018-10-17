package models

import "encoding/json"

type DBServiceInstance struct {
	ID     int             `db:"id"`
	PlanID int             `db:"plan_id"`
	Ctx    json.RawMessage `db:"context"`
	Params json.RawMessage `db:"parameters"`
}
