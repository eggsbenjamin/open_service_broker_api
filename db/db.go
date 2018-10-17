//go:generate mockgen -package db -source=db.go -destination db_mock.go

package db

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// DB defines the database interface.
type DB interface {
	Get(interface{}, string, ...interface{}) error
	Select(interface{}, string, ...interface{}) error
	Query(string, ...interface{}) (Rows, error)
	NamedQuery(string, interface{}) (Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// Beginner defines the db beginnner interface (wraps sqlx for mocking).
type Beginner interface {
	Begin() (DBTxer, error)
}

// Txer defines the transaction interface (wraps sqlx for mocking).
type Txer interface {
	Commit() error
	Rollback() error
}

// DBBeginner defines the db beginnner interface (wraps sqlx for mocking).
type DBBeginner interface {
	DB
	Beginner
}

// DBTxer defines the db transactioner interface (wraps sqlx for mocking).
type DBTxer interface {
	DB
	Txer
}

// Rows defines the rows interface.
type Rows interface {
	Next() bool
	Scan(...interface{}) error
	StructScan(interface{}) error
	Err() error
	Close() error
}

// db implements DB (wraps sqlx to afford mocking)
type db struct {
	*sqlx.DB
}

// NewDB creates instantiates a db object.
func NewDB(sqlxDB *sqlx.DB) DBBeginner {
	return &db{sqlxDB}
}

func (d *db) Get(out interface{}, query string, args ...interface{}) error {
	return d.DB.Get(out, query, args...)
}

func (d *db) Select(out interface{}, query string, args ...interface{}) error {
	return d.DB.Select(out, query, args...)
}

func (d *db) Query(query string, args ...interface{}) (Rows, error) {
	return d.DB.Queryx(query, args...)
}

// NamedQuery wraps the sqlx method of the same name.
func (d *db) NamedQuery(query string, arg interface{}) (Rows, error) {
	return d.DB.NamedQuery(query, arg)
}

// Begin wraps the sqlx method of the same name.
func (d *db) Begin() (DBTxer, error) {
	tx, err := d.DB.Beginx()
	return NewTx(tx), err
}

type tx struct {
	*sqlx.Tx
}

// NewTx creates a new DBTxer
func NewTx(sqlxTx *sqlx.Tx) DBTxer {
	return &tx{sqlxTx}
}

// Query wraps the sqlx method of the same name.
func (t *tx) Query(query string, args ...interface{}) (Rows, error) {
	return t.Tx.Queryx(query, args...)
}

// NamedQuery wraps the sqlx method of the same name.
func (t *tx) NamedQuery(query string, arg interface{}) (Rows, error) {
	return t.Tx.NamedQuery(query, arg)
}

// Exec wraps the sqlx method of the same name.
func (t *tx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return t.Tx.Exec(query, args...) // nolint: safesql
}

// Commit wraps the sqlx method of the same name.
func (t *tx) Commit() error {
	return t.Tx.Commit()
}

// Rollback wraps the sqlx method of the same name.
func (t *tx) Rollback() error {
	return t.Tx.Rollback()
}

// NewConnection returns a new connection to a db
func NewConnection(host, port, user, pwd, name string) (DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		pwd,
		name,
	)
	conn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, errors.Wrapf(err, "NewConnection: error connecting to db")
	}

	return NewDB(conn), nil
}
