package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user ok"))
}
