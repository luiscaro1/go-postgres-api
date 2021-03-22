package utils

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func Router_use(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}
