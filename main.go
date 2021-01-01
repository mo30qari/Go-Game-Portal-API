package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Avatar string `json:"avatar"`
	Password string `json:"password"`
}

type Game struct {
	Name string `json:"username"`
	Genre Genre `json:"genre"`
	Icon string `json:"icon"`
	Link string `json:"link"`
	Owner User `json:"owner"`
}

type Genre struct{
	Name string `json:"name"`
}

func main(){
	//must be deleted
	i, _ := json.MarshalIndent(Genre{Name: "Shooter"}, "", " ")
	fmt.Println(string(i))
	//
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/register/{username}/{password}/{confirm}/{Avatar}", register).Methods("POST")

	http.ListenAndServe(":8080", router)



}
