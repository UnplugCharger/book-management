package routes

import (
	"UnplugCharger/book-management/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/",controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.GetBook).Methods("Get")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
}

