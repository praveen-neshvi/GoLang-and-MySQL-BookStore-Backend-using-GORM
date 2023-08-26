package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/praveen-neshvi/bookstore/pkg/routes"
	// "gorm.io/driver/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisteredBookstoreRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
