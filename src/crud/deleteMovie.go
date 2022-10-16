package crud

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"example-rest-api/src/handler"
	"example-rest-api/src/setup"
	st "example-rest-api/src/structure"
)

// Delete a movie

// response and request handlers
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	movieID := params["movieid"]

	var response = st.JsonResponse{}

	if movieID == "" {
		response = st.JsonResponse{Type: "error", Message: "You are missing movieID parameter."}
	} else {
		db := setup.SetupDB()

		handler.PrintMessage("Deleting movie from DB")

		_, err := db.Exec("DELETE FROM movies where movieID = $1", movieID)

		// check errors
		handler.CheckErr(err)

		response = st.JsonResponse{Type: "success", Message: "The movie has been deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}
