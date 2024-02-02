package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("MyEsportsApp is healthy! Version: %s", GetVersion())
	w.Write([]byte(message))
}

func GetTeamsHandler(w http.ResponseWriter, r *http.Request) {
	teams := teams.GetTeams()
	respondWithJSON(w, http.StatusOK, teams)
}

// GetPlayerHandler returns a JSON profile of a specific player
func GetPlayerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamName := vars["teamName"]
	playerName := vars["playerName"]

	player, found := teams.GetPlayerByName(teamName, playerName)
	if !found {
		respondWithError(w, http.StatusNotFound, "Player not found")
		return
	}

	respondWithJSON(w, http.StatusOK, player)
}

func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}

// respondWithError sends an error response with the specified status code and message
func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	respondWithJSON(w, statusCode, map[string]string{"error": message})
}
