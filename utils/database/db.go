package database

import (
	"database/sql"
)

// This is a basic interface for self-use database.
// Just in case that it might be changed to use different types of databases,
// we can do it with out affecting the upper level
type Database interface {

	// Selects one item, result will be stored in "out". It will be casted
	SelectOne(out interface{}, sql string, args ...interface{}) error

	// Selects multiple items, result will be stored in "out". It will be casted
	SelectMany(out interface{}, sql string, args ...interface{}) error

	// Queries for "SELECT <x> WHERE <y> in (...)". Got results in "out" interface (with casting)
	SelectIn(out interface{}, sql string, args ...interface{}) error

	// Queries for "SELECT <x> WHERE <y> in (...)". Got results in returning rows
	QueryIn(sql string, args ...interface{}) (*sql.Rows, error)

	// Executes an sql request, usually do not expect returning rows
	Exec(sql string, args ...interface{}) (sql.Result, error)

	// Executes an sql request, expect returning data (rows)
	ExecReturning(sql string, args ...interface{}) (*sql.Rows, error)
}
