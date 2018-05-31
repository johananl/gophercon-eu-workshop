package routing

import (
	"net/http"

	"github.com/gorilla/mux"
)

func DiagnosticsRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/healthz", readyzHandler()).Methods(http.MethodGet)
	r.HandleFunc("/readyz", healthzHandler()).Methods(http.MethodGet)

	return r
}
