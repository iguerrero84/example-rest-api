package crud

import (
	"encoding/json"
	"net/http"

	"example-rest-api/src/handler"
	"example-rest-api/src/setup"
	st "example-rest-api/src/structure"
)

// Get all movies

// response and request handlers
func GetMovies(w http.ResponseWriter, r *http.Request) {
	db := setup.SetupDB()

	handler.PrintMessage("Getting movies...")

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM movies")

	// check errors
	handler.CheckErr(err)

	// var response []JsonResponse
	var movies []st.Movie

	// Foreach movie
	for rows.Next() {
		var id int
		var movieID string
		var movieName string

		err = rows.Scan(&id, &movieID, &movieName)

		// check errors
		handler.CheckErr(err)

		movies = append(movies, st.Movie{MovieID: movieID, MovieName: movieName})
	}

	var response = st.JsonResponse{Type: "success", Data: movies}

	json.NewEncoder(w).Encode(response)
}
