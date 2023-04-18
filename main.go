package main

import (
	"encoding/json"
	"net/http"
)

func main(){
	http.HandleFunc("/person/", func(w http.ResponseWriter, r *http.Request){
		if r.Method =="POST" {
			// Criação da pessoa
			type Person struct {
				Username string 'json: "username"'
				Password string 'json: "password"'
				Full name string 'json: "full name"'
			}
			var person Person
			err := json.NewDecoder(r.Body).Decode(&person)
			if err != nil {
				http.Error(w, "Error trying to create person", http.StatusBadRequest)
			}
		}
		http.Error(w, "Not implemented", http.StatusInternalServerError) 
	})


	http.ListenAndServe(":8080", nil)
}

