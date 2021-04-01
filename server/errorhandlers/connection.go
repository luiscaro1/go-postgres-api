package errorhandlers

import (
	"net/http"
)

// CheckError verifies an incoming error and sends a bad request error
func CheckError(err error, w http.ResponseWriter) {
	if err != nil && w != nil {
		w.WriteHeader(http.StatusBadRequest)

	}
}
