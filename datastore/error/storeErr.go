package error

import (
	"database/sql"
	"errors"
)

var (
	// ErrUnrecognizedServiceModel is the error when service model cannot be recognized
	ErrUnrecognizedServiceModel = errors.New("unrecognized service model")
	// ErrUnrecognizedDBModel is the error when DB model cannot be recognized
	ErrUnrecognizedDBModel = errors.New("unrecognized database model")
	// ErrConverterError is the error when models cannot be converted
	ErrConverterError = errors.New("converter did not work correctly")
	// ErrInvalidInputModel is the error when input model is not
	ErrInvalidInputModel = errors.New("input model type is not correct")
	// ErrDBRetrievalError is the error when db retrieval is not successful
	ErrDBRetrievalError = errors.New("failed to retrieve from db")
	// ErrInvalidQuery is the error when query is invalid
	ErrInvalidQuery = errors.New("query not valid")
	// ErrNotSupportedQuery is the error when query is not supported
	ErrNotSupportedQuery = errors.New("query not supported yet")
	// ErrNotFound is the error when data is not found
	ErrNotFound = errors.New("not found")
)

// Wrapper wraps some db errors with store errors for outside use
func Wrapper(err error) error {
	switch {
	case err == sql.ErrNoRows:
		return ErrNotFound
	default:
		return ErrDBRetrievalError
	}
}
