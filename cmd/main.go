package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/carlosfgti/go-jwt-postgresql/cmd/controllers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env: %v", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/health", healthHandler).Methods("GET")
	router.HandleFunc("/api/v1/users", controllers.CreateUser).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Listening on port %s\n", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
