package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mvzcanhaco/api-verifymy-crud-test/controllers"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/users", controllers.AllUsers).Methods("Get")
	r.HandleFunc("/api/users/{id}", controllers.ReturnUser).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
