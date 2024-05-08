package main

import (
	"log"
	"net/http"

	"github.com/batyrbek/pkg/handler"
	"github.com/batyrbek/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	r.HandleFunc("/register", handler.SignUp).Methods("POST")
	r.HandleFunc("/login", handler.Login).Methods("POST")
	r.HandleFunc("/logout", handler.Logout).Methods("GET")

	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
