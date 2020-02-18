package ping

import (
	"net/http"
	"encoding/json"
)

const (
	PONG = "pong"
)

func PingRouteHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(PONG)
}
