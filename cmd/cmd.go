package cmd

import (
	"Advanced_programming_project/db"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RunServer() {
	err := db.DbConnection()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	router := mux.NewRouter()
	setupRoutes(router)

	port := ":3001"
	fmt.Printf("Starting server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
