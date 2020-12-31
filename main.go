package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type User struct {
	Username string
	Avatar string
	Password string
}

type Game struct {
	Name string
	Genre Genre
	Icon string
	Link string
	Owner User
}

type Genre struct{
	Name string
}

func main(){

	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/register/{username}/{password}/{confirm}/{Avatar}", register).Methods("POST")

	http.ListenAndServe(":8080", router)

}
