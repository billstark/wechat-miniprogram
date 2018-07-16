package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"

	// for postgres driver type
	_ "github.com/lib/pq"
)

const (
	poetgres                  = "postgres"
	doesNotSupport            = "database: PsqlDB cannot initialize other types of DB."
	connectionDefaultTemplate = "host=%s port=%s dbname=%s user=%s password='%s'"
	connectionGAETemplate     = "host=%s dbname=%s user=%s password='%s'"
	connectionNoSSL           = " sslmode=disable"
)

// PsqlDB create a new db struct
type PsqlDB struct {
	connection *sqlx.DB
}

// New Creates a new psql database (most basic and traditional way)
func New(config DBConfig) (Database, error) {

	// Only allows postgres
	if config.Driver != poetgres {
		return nil, errors.New(doesNotSupport)
	}

	// Gets connection string
	connectionString := fmt.Sprintf(connectionDefaultTemplate,
		config.Host, strconv.Itoa(config.Port), config.DBName, config.Username, config.Password,
	)

	if !config.EnableSSL {
		connectionString += connectionNoSSL
	}

	// Makes connection
	connection, err := sqlx.Connect(poetgres, connectionString)
	if err != nil {
		return nil, err
	}

	connection.SetMaxIdleConns(config.MaxIdleConnections)
	connection.SetMaxOpenConns(config.MaxOpenConnections)

	return &PsqlDB{connection}, nil
}

// NewGAE creates a new psql database (for Google App Engine only)
func NewGAE(config DBConfig) (Database, error) {

	// Only allows postgres
	if config.Driver != poetgres {
		return nil, errors.New(doesNotSupport)
	}

	connectionString := fmt.Sprintf(connectionGAETemplate,
		config.Host, config.DBName, config.Username, config.Password,
	)

	if !config.EnableSSL {
		connectionString += connectionNoSSL
	}

	// Makes connection
	connection, err := sqlx.Open(poetgres, connectionString)
	if err != nil {
		return nil, err
	}

	connection.SetMaxIdleConns(config.MaxIdleConnections)
	connection.SetMaxOpenConns(config.MaxOpenConnections)

	return &PsqlDB{connection}, nil
}

// SelectOne selects one item. result in "out" (casted)
func (d *PsqlDB) SelectOne(out interface{}, sql string, args ...interface{}) error {
	return d.connection.Get(out, sql, args...)
}

// SelectMany selects several items. result in "out" (casted)
func (d *PsqlDB) SelectMany(out interface{}, sql string, args ...interface{}) error {
	return d.connection.Select(out, sql, args...)
}

// SelectIn executes "SELECT <x> WHERE <y> in (...)", result in "out" (casted)
func (d *PsqlDB) SelectIn(out interface{}, sql string, args ...interface{}) error {
	query, args, err := sqlx.In(sql, args...)
	if err != nil {
		return err
	}
	query = d.connection.Rebind(query)
	return d.connection.Select(out, query, args...)
}

// QueryIn executes "SELECT <x> WHERE <y> in (...)", result will be returned as sql rows
func (d *PsqlDB) QueryIn(sql string, args ...interface{}) (*sql.Rows, error) {
	query, args, err := sqlx.In(sql, args...)
	if err != nil {
		return nil, err
	}
	query = d.connection.Rebind(query)
	return d.connection.Query(query, args...)
}

// Exec executes query, usually non return values
func (d *PsqlDB) Exec(sql string, args ...interface{}) (sql.Result, error) {
	return d.connection.Exec(sql, args...)
}

// ExecReturning executes query, usually have return values as sql rows
func (d *PsqlDB) ExecReturning(sql string, args ...interface{}) (*sql.Rows, error) {
	return d.connection.Query(sql, args...)
}
