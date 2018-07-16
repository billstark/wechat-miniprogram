package healthcheck

import (
	"encoding/json"
	"net/http"
	"time"
)

const (
	HTTP_HEADER_CONTENT = "Content-Type"
	HTTP_CONTENT_JSON   = "application/json"
	HTTP_CONTENT_UTF8   = "charset=utf-8"
	HTTP_HEADER_BREAK   = ";"

	RESPONSE_OK_TAG   = "ok"
	RESPONSE_TIME_TAG = "time"
)

// Simple is a health check which returns `ok: true` and the current Unix time.
func Simple(w http.ResponseWriter, req *http.Request) {
	w.Header().Set(HTTP_HEADER_CONTENT, HTTP_CONTENT_JSON+HTTP_HEADER_BREAK+HTTP_CONTENT_UTF8)
	json.NewEncoder(w).Encode(&map[string]interface{}{
		RESPONSE_OK_TAG:   true,
		RESPONSE_TIME_TAG: time.Now().Unix(),
	})
}
