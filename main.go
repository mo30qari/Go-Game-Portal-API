package main

import (
	_ "encoding/json"
	_ "fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type User struct {
	Username string `json:"username" validate:"string,Min=3,Max=9"`
	Email string `json:"email" validate:"email"`
	//Avatar string `json:"avatar" validate:"file,format=png|jpg|jpeg"`
	Password string `json:"password" validate:"password,Min=5,Max=255"`
}

//Incomplete
type Game struct {
	Name string `json:"username"`
	Genre Genre `json:"genre"`
	Icon string `json:"icon"`
	Link string `json:"link"`
	Owner User `json:"owner"`
}

//Incomplete
type Genre struct{
	Name string `json:"name"`
}

func main(){

	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/register/{username}/{email}/{password}", register).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))

}