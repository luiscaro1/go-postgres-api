package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luiscaro1/go-postgres-api/server/routes"
	"github.com/luiscaro1/go-postgres-api/server/utils"
)

func main() {

	router := mux.NewRouter()
	utils.RouterUse(router, "/api/user", routes.UserRoutes())

	log.Fatal(http.ListenAndServe(":8000", router))

}
