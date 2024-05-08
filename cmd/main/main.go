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

	r.HandleFunc("/register", handler.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", handler.LoginHandler).Methods("POST")

	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
