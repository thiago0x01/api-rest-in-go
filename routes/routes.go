package routes

import (
	"api-rest-in-go/controllers"
	"api-rest-in-go/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.Use(middlewares.SetContentType)

	r.HandleFunc("/", controllers.Home)

	r.HandleFunc("/personalities", controllers.ReadPersonalities).Methods("Get")
	r.HandleFunc("/personalities/{id}", controllers.ReadPersonalityById).Methods("Get")
	r.HandleFunc("/personalities", controllers.CreatePersonality).Methods("Post")
	r.HandleFunc("/personalities/{id}", controllers.DeletePersonality).Methods("Delete")
	r.HandleFunc("/personalities/{id}", controllers.UpdatePersonality).Methods("Put")

	log.Fatal(http.ListenAndServe(":5000", r), handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r))
}
