package crud

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example-rest-api/src/handler"
	"example-rest-api/src/setup"
	st "example-rest-api/src/structure"
)

// Create a movie

// response and request handlers
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	movieID := r.FormValue("movieid")
	movieName := r.FormValue("moviename")

	var response = st.JsonResponse{}

	if movieID == "" || movieName == "" {
		response = st.JsonResponse{Type: "error", Message: "You are missing movieID or movieName parameter."}
	} else {
		db := setup.SetupDB()

		handler.PrintMessage("Inserting movie into DB")

		fmt.Println("Inserting new movie with ID: " + movieID + " and name: " + movieName)

		var lastInsertID int
		err := db.QueryRow("INSERT INTO movies(movieID, movieName) VALUES($1, $2) returning id;", movieID, movieName).Scan(&lastInsertID)

		// check errors
		handler.CheckErr(err)

		response = st.JsonResponse{Type: "success", Message: "The movie has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}
