package testutils

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"wechat-miniprogram/application"
)

// CheckResponseCode checks response code equal
func CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

// ExecuteRequest executes a http request
func ExecuteRequest(req *http.Request, app application.App) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)
	return rr
}
