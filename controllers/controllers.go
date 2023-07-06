package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mvzcanhaco/api-verifymy-crud-test/database"
	"github.com/mvzcanhaco/api-verifymy-crud-test/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	var p []models.User
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func ReturnUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user models.User
	database.DB.First(&user, id)
	json.NewEncoder(w).Encode(user)
}

func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)
	database.DB.Create(&newUser)
	json.NewEncoder(w).Encode(newUser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var user models.User
	database.DB.Delete(&user, id)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var user models.User
	database.DB.First(&user, id)
	json.NewDecoder(r.Body).Decode(&user)
	database.DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}
