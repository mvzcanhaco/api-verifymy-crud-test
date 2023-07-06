package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mvzcanhaco/api-verifymy-crud-test/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Users)
}

func ReturnUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, User := range models.Users {
		if strconv.Itoa(int(User.Id)) == id {
			json.NewEncoder(w).Encode(User)
		}
	}
}
