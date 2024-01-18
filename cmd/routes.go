package cmd

import (
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes(router *mux.Router) {
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("frontend/static"))))

	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/teach/{id}", volunteerPersonalPageHandler).Methods("GET")
	router.HandleFunc("/stud/{id}", childPersonalPageHandler).Methods("GET")
	router.HandleFunc("/teachlogin", teachLoginHandler).Methods("GET", "POST")
	router.HandleFunc("/teachreg", teachRegHandler).Methods("GET", "POST")
	router.HandleFunc("/studlog", studLogHandler).Methods("GET", "POST")
	router.HandleFunc("/studreg", studRegHandler).Methods("GET", "POST")
}
