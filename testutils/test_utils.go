package testutils

import (
	"github.com/eggsbenjamin/open_service_broker_api/db"
	"github.com/pkg/errors"
)

// Empties all tables in db
func TeardownDB(db db.DB) error {
	// note: order of table deletion is important
	tables := []string{
		"service_instance",
		"service_plan",
		"service",
	}

	for _, table := range tables {
		if _, err := db.Exec("DELETE FROM " + table); err != nil {
			return errors.Wrapf(err, "TeardownDB: error deleting from table: %s", table)
		}
	}

	return nil
}
