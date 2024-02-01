package cmd

import (
	"net/http"

	"github.com/gorilla/mux"	
)

func setupRoutes(router *mux.Router) {
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("frontend/static"))))

	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/teacher/{_id}", teacherPersonalPageHandler).Methods("GET")
	router.HandleFunc("/student/{_id}", studentPersonalPageHandler).Methods("GET")
	router.HandleFunc("/teacherlogin", teachLoginHandler).Methods("GET", "POST")
	router.HandleFunc("/teacherreg", teachRegHandler).Methods("GET", "POST")
	router.HandleFunc("/studentlogin", studLogHandler).Methods("GET", "POST")
	router.HandleFunc("/studentreg", studRegHandler).Methods("GET", "POST")
}
