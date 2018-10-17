package models

type DBServicePlan struct {
	ID        int    `db:"id"`
	ServiceID int    `db:"service_id"`
	Name      string `db:"name"`
}
