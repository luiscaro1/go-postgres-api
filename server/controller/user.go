package controller

/*
	*** We can pass a map to the decode funcion to get a MAP instead of struct ***

	*** Something like this ***
	var userMap map[string]string
*/

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/luiscaro1/go-postgres-api/server/errorhandlers"
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

	// Parse the URI for paramaters
	u, err := url.ParseRequestURI(r.RequestURI)

	// get a map of the query
	queries := u.Query()

	// establish a connection with the db
	db := services.OpenConnection()

	// verify if an error ocurred during parsing the params
	errorhandlers.CheckError(err, w)

	// Initializing the query string for the GET request
	var queryString string = "SELECT * FROM mock_data"

	// verify if the URI has any parameters, if so add the WHERE clause for selection
	if len(queries) > 0 {
		queryString += "\t WHERE"
	}

	/*
		arguments (slice) that are going to be passed to the query function
		since no specific datatype can be assigned the interface values is chosen
	*/
	var args []interface{}

	// number used as placeholder in the query string
	placeHolderIdx := 1

	/*
		loop through fields in the body and append them to the query selector
		 as well take the values and append them to the args slice for later passing
	*/

	for k, v := range queries {

		queryString += "\t" + k + "=$" + strconv.Itoa(placeHolderIdx)

		// while we still have arguments we want to add AND as long as the argument isn't the last in the slice
		if placeHolderIdx != len(queries) {
			queryString += "\tAND\t"
		}

		// increment the index for the next placeholder
		placeHolderIdx++

		// add the body values to the args slice
		args = append(args, v[0])
	}

	/*
	 execute the query
	 we use the unpack operator to unpack all of the arguments
	*/
	rows, err := db.Query(queryString, args...)
	fmt.Println(rows, "here")

	// verify if an error has ocurred while making the q
	errorhandlers.CheckError(err, w)

	// creat a slice of user structs
	var users []User

	// iterate through through the Users table by rows
	for rows.Next() {

		// temp user
		var user User

		// grab the values in the rows and place them in the fields memory addresses
		err = rows.Scan(&user.FirstName, &user.LastName, &user.Email)

		// append all the Users structs to the slice
		users = append(users, user)

		// verify if any erros have occured while scanning
		errorhandlers.CheckError(err, w)

	}
	// convert json to an array of byes to write to the response
	userBytes, _ := json.Marshal(users)

	// set the headers for a JSON response
	w.Header().Set("Content-Type", "application/json")

	// set stus to ok if no error appear
	w.WriteHeader(http.StatusOK)

	// send the stream of bytes to the client
	w.Write(userBytes)

	// close connection to the table  & db once response is sent
	defer rows.Close()
	defer db.Close()

}

// UserPost adds a new user to the db
func UserPost(w http.ResponseWriter, r *http.Request) {

	// open connection to database for requests
	db := services.OpenConnection()

	// create a new user instance
	var user User

	// destructure the body into the user object
	err := json.NewDecoder(r.Body).Decode(&user)

	// verify if any errors arise
	errorhandlers.CheckError(err, w)

	// Inset a new User
	_, err = db.Exec(`INSERT INTO mock_data (first_name, last_name, email) VALUES ($1, $2, $3)`, user.FirstName, user.LastName, user.Email)

	// verify if the query produces any errors
	errorhandlers.CheckError(err, w)

	//close connectiion to the database once the response is sent
	defer db.Close()

	// send code 201 to show something was created
	w.WriteHeader(http.StatusCreated)

}
