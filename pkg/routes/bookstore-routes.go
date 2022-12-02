package routes

import (
	"github.com/gorilla/mux"
	"github.com/mikespinks0401/go-bookManagement/pkg/controllers"
)

var RegisterBookManagementeRoutes = func(router *mux.Router){
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId", controllers.DeleteBook).Methods("DELETE")
}