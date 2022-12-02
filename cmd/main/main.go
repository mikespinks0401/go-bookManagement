package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mikespinks0401/go-bookManagement/pkg/routes"
)



func main(){
	r := mux.NewRouter()
	routes.RegisterBookManagementeRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
	
}