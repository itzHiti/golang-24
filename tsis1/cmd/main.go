package main

import (
	"net/http"

	handlers "tsis1/pkg"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/teams", handlers.GetTeamsHandler).Methods("GET")
	router.HandleFunc("/teams/{teamName}/players/{playerName}", handlers.GetPlayerHandler).Methods("GET")
	router.HandleFunc("/health", handlers.HealthCheck).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}
