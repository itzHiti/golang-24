package cmd

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}
