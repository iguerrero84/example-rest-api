package crud

import (
	"encoding/json"
	"net/http"

	"example-rest-api/src/handler"
	"example-rest-api/src/setup"
	st "example-rest-api/src/structure"
)

// Delete all movies

// response and request handlers
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	db := setup.SetupDB()

	handler.PrintMessage("Deleting all movies...")

	_, err := db.Exec("DELETE FROM movies")

	// check errors
	handler.CheckErr(err)

	handler.PrintMessage("All movies have been deleted successfully!")

	var response = st.JsonResponse{Type: "success", Message: "All movies have been deleted successfully!"}

	json.NewEncoder(w).Encode(response)
}
