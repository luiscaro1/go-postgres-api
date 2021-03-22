package routes

import (
	"github.com/gorilla/mux"
	"github.com/luiscaro1/go-postgres-api/server/controller"
)

func UserRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/all", controller.UserGet).Methods("GET")
	return router
}
