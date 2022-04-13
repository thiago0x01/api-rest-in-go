package controllers

import (
	"api-rest-in-go/database"
	"api-rest-in-go/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func ReadPersons(w http.ResponseWriter, r *http.Request) {
	var person []models.Person

	database.DB.Find(&person)

	json.NewEncoder(w).Encode(person)
}

func ReadPersonById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var person models.Person

	database.DB.First(&person, id)

	json.NewEncoder(w).Encode(person)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person

	json.NewDecoder(r.Body).Decode(&person)

	database.DB.Create(&person)

	json.NewEncoder(w).Encode(person)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var person models.Person

	database.DB.Delete(&person, id)

	json.NewEncoder(w).Encode(person)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var person models.Person

	database.DB.First(&person, id)

	json.NewDecoder(r.Body).Decode(&person)

	database.DB.Save(&person)

	json.NewEncoder(w).Encode(person)
}
