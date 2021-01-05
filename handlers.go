package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode("Welcome to My Game Portal! You see nothing here. But you can manage your games if you know how to work with my API. I'm gonna to add the API documentation in it's Github page.")

}

func register(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	user := User{
		Username: params["username"],
		Email: params["email"],
		Password: params["password"],
	}

	for i, err := range validateStruct(user) {
		fmt.Printf("%d: %s", i+1, err.Error())
	}

}
