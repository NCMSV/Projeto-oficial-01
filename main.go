package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NCMSV/Projeto-oficial-01/domain"
	"github.com/NCMSV/Projeto-oficial-01/domain/person"
)

func main() {
	personService, err := person.NewService("person,json")
	if err != nil {
		fmt.Println("Error trying to create person service")
		return
	}

	http.HandleFunc("/person/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

			// Criação da pessoa
			var person domain.Person
			err := json.NewDecoder(r.Body).Decode(&person)
			if err != nil {
				fmt.Printf("Error trying to decode body. Body Should be a json. Error: %s", err.Error())
				http.Error(w, "Error trying to create person", http.StatusBadRequest)
				return
			}
			if person.ID <= 0 {
				http.Error(w, "Error trying to create person. Username should be a positive integer", http.StatusBadRequest)
				return
			}
			///Criar pessoa
			err = personService.Create(person)
			if err != nil {
				fmt.Printf("Error trying to create person: %s", err.Error())
				http.Error(w, "Error trying to create person", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
			return
		}
		http.Error(w, "Not implemented", http.StatusInternalServerError)
	})

	http.ListenAndServe(":8080", nil)
}
