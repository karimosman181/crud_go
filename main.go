package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"diretor"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

/**
 * gets all the movies
**/
func getMovies(w http.ResponseWriter, r *http.Request) {
	//setting header content type to json
	w.Header().Set("content-Type", "application/json")

	//encodig reponse and movies to json and return
	json.NewEncoder(w).Encode(movies)
}

/**
 * gets a movie
**/
func getMovie(w http.ResponseWriter, r *http.Request) {
	//setting header content type to json
	w.Header().Set("content-Type", "application/json")

	//getting params from header
	params := mux.Vars(r)

	// loop over movies
	for _, item := range movies {

		//check if item id equals params id
		if item.ID == params["id"] {

			//encodig reponse and item to json and return
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

/**
 * create a movie
**/
func createMovie(w http.ResponseWriter, r *http.Request) {
	//setting header content type to json
	w.Header().Set("content-Type", "application/json")

	//declare a movie
	var movie Movie

	//decode body and refernce it to movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	//create random id
	movie.ID = strconv.Itoa(rand.Intn(10000000000))

	//add item to movies
	movies = append(movies, movie)

	//encodig reponse and movie to json and return
	json.NewEncoder(w).Encode(movie)
}

/**
 * update a movie
**/
func updateMovie(w http.ResponseWriter, r *http.Request) {
	//setting header content type to json
	w.Header().Set("content-Type", "application/json")

	//getting params from header
	params := mux.Vars(r)

	//declare a movie
	var movie Movie

	//decode body and refernce it to movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	//set movie id
	movie.ID = params["id"]

	// loop over movies
	for index, item := range movies {

		//check if item id equals params id
		if item.ID == params["id"] {
			//remove the old movie than add the new movie in the same place

			//remove the old movie
			movies = append(movies[:index], movies[index+1:]...)

			//add new movie
			movies = append(movies, movie)

			//encodig reponse and movie to json and return
			json.NewEncoder(w).Encode(movie)
			return
		}

	}
}

/**
 * delete a movie
**/
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	//setting header content type to json
	w.Header().Set("content-Type", "application/json")

	//getting params from header
	params := mux.Vars(r)

	// loop over movies
	for index, item := range movies {

		//check if item id equals params id
		if item.ID == params["id"] {

			//remove the item by appending the reset of the list in the item place
			movies = append(movies[:index], movies[index+1:]...)
			break
		}

	}

	//encodig reponse and movies to json and return the rest of the movies
	json.NewEncoder(w).Encode(movies)
}

/**
 * main function
**/
func main() {
	//using gorilla mux for routing
	r := mux.NewRouter()

	//adding items to movies
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "438228", Title: "Movie Two", Director: &Director{Firstname: "Tom", Lastname: "Something"}})

	//creating routes
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	//serving at port 8000
	fmt.Printf("Staring server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
