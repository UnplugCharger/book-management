package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/UnplugCharger/book-management/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	http.Handle("/", r)
	fmt.Println("starting server at  localhost:9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
