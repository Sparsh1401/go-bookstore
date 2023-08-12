package main

import (
	"net/http"

	"github.com/Sparsh1401/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	myRouter := mux.NewRouter()

	routes.RegisterBookStoreRoutes(myRouter)

	http.Handle("/", myRouter)

	http.ListenAndServe(":80", myRouter)
}
