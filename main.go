package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Rating    string    `json:"rating"`
	Directior *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["ID"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Title: "Bahubali", Rating: "10", Directior: &Director{Firstname: "SS", Lastname: "Rajamoule"}})
	movies = append(movies, Movie{ID: "2", Title: "Liger", Rating: "8", Directior: &Director{Firstname: "PURI", Lastname: "Jagannath"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{ID}", getMovie).Methods("GET")

	fmt.Println("Starting server on 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
