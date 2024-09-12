package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/carlosfgti/go-jwt-postgresql/cmd/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if user.Name == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "Missing fields", http.StatusBadRequest)
		return
	}
	token := user.Create(map[string]interface{}{})
	w.Write([]byte(token))
	w.WriteHeader(http.StatusCreated)
}
