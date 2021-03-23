package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/luiscaro1/go-postgres-api/server/error_handlers"
	"github.com/luiscaro1/go-postgres-api/server/services"
)

// User holds the formatting for users in the db
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// UserGet returns all of the users in the db
func UserGet(w http.ResponseWriter, r *http.Request) {
	db := services.OpenConnection()

	rows, err := db.Query("SELECT * FROM mock_data")

	error_handlers.CheckError(err, w)

	var users []User

	for rows.Next() {
		var user User

		err = rows.Scan(&user.FirstName, &user.LastName, &user.Email)
		users = append(users, user)
		fmt.Println(user)
		error_handlers.CheckError(err, w)

	}

	userBytes, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userBytes)

	defer rows.Close()
	defer db.Close()

}

// UserPost adds a new user to the db
func UserPost(w http.ResponseWriter, r *http.Request) {

	db := services.OpenConnection()

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)

	error_handlers.CheckError(err, w)
	_, err = db.Exec(`INSERT INTO mock_data (first_name, last_name, email) VALUES ($1, $2, $3)`, user.FirstName, user.LastName, user.Email)
	error_handlers.CheckError(err, w)

	defer db.Close()

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(user)

}
