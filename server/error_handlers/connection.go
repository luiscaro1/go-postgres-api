package error_handlers

import "net/http"

func CheckError(err error, w http.ResponseWriter) {
	if err != nil {
		if w != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		panic(err)
	}
}
