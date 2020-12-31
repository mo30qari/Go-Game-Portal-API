package main

import (
	"encoding/json"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){

	json.NewEncoder(w).Encode("Welcome to My Game Portal! You see nothing here. But you can manage your games if you know how to work with my API. I'm gonna to add the API documentation in it's Github page.")

}

func register(w http.ResponseWriter, r *http.Request){

	openDb()

}
