package routes

import (
	"github.com/febriyantoilham/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router)  {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
}