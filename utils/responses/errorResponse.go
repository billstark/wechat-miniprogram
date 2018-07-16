package responses

type ErrorResponse struct {
	Status  int      `json:"-"`
	Message string   `json:"error"`
	Details []string `json:"details"`
}

const (
	STATUS_INVALID           = 400
	STATUS_PERMISSION_DENIED = 401
	STATUS_NOTFOUND          = 404
	STATUS_INERNAL_ERROR     = 500

	MESSAGE_ACCESS_DENIED  = "access denied"
	MESSAGE_NOT_FOUND      = "not found"
	MESSAGE_UNAUTHORIZED   = "unauthorized"
	MESSAGE_INTERNAL_ERROR = "internal server error"
)

// Invalid creates an ErrorResponse for an invalid request.
func Invalid(message string, details []string) ErrorResponse {
	return ErrorResponse{STATUS_INVALID, message, details}
}

// Denied creates a 401 Unauthorized response.
func Denied() ErrorResponse {
	return ErrorResponse{STATUS_PERMISSION_DENIED, MESSAGE_ACCESS_DENIED, nil}
}

// NotFound creates a 404 Not Found response.
func NotFound() ErrorResponse {
	return ErrorResponse{STATUS_NOTFOUND, MESSAGE_NOT_FOUND, nil}
}

// Unauthorized creates a 401 Unauthorized response with an optional message.
func Unauthorized(message string) ErrorResponse {
	if len(message) == 0 {
		message = MESSAGE_UNAUTHORIZED
	}
	return ErrorResponse{STATUS_PERMISSION_DENIED, message, nil}
}

// InternalError creates a 500 Internal Server Error response with an optional message.
func InternalError(message string) ErrorResponse {
	if len(message) == 0 {
		message = MESSAGE_INTERNAL_ERROR
	}
	return ErrorResponse{STATUS_INERNAL_ERROR, message, nil}
}
