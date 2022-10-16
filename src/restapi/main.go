package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"example-rest-api/src/crud"
	st "example-rest-api/src/structure"
)

type JsonResponse struct {
	Type    string     `json:"type"`
	Data    []st.Movie `json:"data"`
	Message string     `json:"message"`
}

// Main function
func main() {

	//get the value of the ADDR environment variable
	addr := os.Getenv("ADDR")

	//if it's blank, default to ":80", which means
	//listen port 80 for requests addressed to any host
	if len(addr) == 0 {
		addr = ":80"
	}

	// Init the mux router
	router := mux.NewRouter()

	// Route handles & endpoints

	// Get all movies
	router.HandleFunc("/movies/", crud.GetMovies).Methods("GET")

	// Create a movie
	router.HandleFunc("/new_movie/", crud.CreateMovie).Methods("POST")

	// Delete a specific movie by the movieID
	router.HandleFunc("/movies/{movieid}", crud.DeleteMovie).Methods("DELETE")

	// Delete all movies
	router.HandleFunc("/movies/", crud.DeleteAllMovies).Methods("DELETE")

	//start the web server using the mux as the root handler,
	//and report any errors that occur.
	//the ListenAndServe() function will block so
	//this program will continue to run until killed
	log.Printf("server is listening at %s...", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
