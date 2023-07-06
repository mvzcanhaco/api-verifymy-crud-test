package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mvzcanhaco/api-verifymy-crud-test/controllers"
	"github.com/mvzcanhaco/api-verifymy-crud-test/middleware"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/users", controllers.AllUsers).Methods("Get")
	r.HandleFunc("/api/users/{id}", controllers.ReturnUser).Methods("Get")
	r.HandleFunc("/api/users", controllers.CreateNewUser).Methods("Post")
	r.HandleFunc("/api/users/{id}", controllers.DeleteUser).Methods("Delete")
	r.HandleFunc("/api/users/{id}", controllers.UpdateUser).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", r))
}
