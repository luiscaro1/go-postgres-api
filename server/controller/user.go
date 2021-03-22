package controller

import (
	"encoding/json"
	"net/http"

	. "github.com/luiscaro1/go-postgres-api/server/error_handlers"
	. "github.com/luiscaro1/go-postgres-api/server/services"
)

type User struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
}

func UserGet(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	rows, err := db.Query("SELECT * FROM mock_data")

	CheckError(err)

	var users []User

	for rows.Next() {
		var user User

		err = rows.Scan(&user.First_name, &user.Last_name, &user.Email)
		users = append(users, user)
		CheckError(err)

	}

	userBytes, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	w.Write(userBytes)

	defer rows.Close()
	defer db.Close()

}
