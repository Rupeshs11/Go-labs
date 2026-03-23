package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/rupeshs11/BookStore_mysql/pkg/models"
	"github.com/rupeshs11/BookStore_mysql/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting BookStore server at http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
