package routes

import (
	"github.com/gorilla/mux"
	"github.com/luiscaro1/go-postgres-api/server/controller"
)

// UserRoutes holds all the routes related to the User
func UserRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/all", controller.UserGet).Methods("GET")
	router.HandleFunc("/add", controller.UserPost).Methods("POST")
	return router
}
