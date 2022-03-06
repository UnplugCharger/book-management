package main

import (
	"UnplugCharger/book-management/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)




func main()  {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:9010",r))
}