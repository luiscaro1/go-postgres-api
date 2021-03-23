package utils

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// RouterUse groups similiar routes with each other and it appends it to the main router
func RouterUse(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}
