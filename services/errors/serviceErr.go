package error

import "errors"

var (
	// ErrIncorrectParamsFormat specifies error that params are not in correct format
	ErrIncorrectParamsFormat = errors.New("incorrect input parameter format")
	// ErrInsufficientParams specifies error that input is not sufficient
	ErrInsufficientParams = errors.New("insufficient parameters")
)
