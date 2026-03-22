package main

import(
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string `json:"id"`
	Isbn     string `json:"isbn"`
	Title    string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content.Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for _,item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}

	}
}

func createMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	_ =json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies =append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter,r *http.Request){
	// set json content type
	w.Header().Set("Content-Type","application/json")
	//params
	params := mux.Vars(r)
	//loop over the movies , range
	//delete the movie with i.d that we have sent
	//add a new movie - the movie that we send in the body of postman
	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			var movie Movie
			_ =json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies =append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}



func main() {
	r := mux.NewRouter()

	movies = append(movies,Movie{ID: "1", Isbn:"2004",Title:"Movie One",Director : &Director{Firstname:"Jethalal",Lastname:"Gada"}})
	movies = append(movies,Movie{ID: "2", Isbn:"2005",Title:"Movie Two",Director : &Director{Firstname:"Champaklal",Lastname:"Gada"}})
	movies = append(movies,Movie{ID: "3", Isbn:"2006",Title:"Movie Three",Director : &Director{Firstname:"Steve",Lastname:"Rogers"}})
	movies = append(movies,Movie{ID: "4", Isbn:"2017",Title:"Movie Four",Director : &Director{Firstname:"Tony",Lastname:"straks"}})
	movies = append(movies,Movie{ID: "5", Isbn:"2019",Title:"Movie Five",Director : &Director{Firstname:"Thor",Lastname:"Odinson"}})
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}
