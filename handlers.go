package main

import (
	"encoding/json"
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
		Email:    params["email"],
		Password: params["password"],
	}

	result := validateStruct(user)
	w.Header().Set("Content-Type", "application/json")

	if len(result) != 0 {
		json.NewEncoder(w).Encode(result)
	} else {
		json.NewEncoder(w).Encode(user)
	}

}
