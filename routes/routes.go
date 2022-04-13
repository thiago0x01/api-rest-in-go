package routes

import (
	"api-rest-in-go/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)

	r.HandleFunc("/persons", controllers.ReadPersons).Methods("Get")
	r.HandleFunc("/persons/{id}", controllers.ReadPersonById).Methods("Get")
	r.HandleFunc("/persons", controllers.CreatePerson).Methods("Post")
	r.HandleFunc("/persons/{id}", controllers.DeletePerson).Methods("Delete")
	r.HandleFunc("/persons/{id}", controllers.UpdatePerson).Methods("Put")

	log.Fatal(http.ListenAndServe(":5000", r))
}
