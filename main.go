package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	db, err := CreateConn()
	if err != nil {
		log.Printf("Failed to create database connection: %v", err)
		return
	}
	// Always close the database when main exits
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		} else {
			log.Println("Database connection closed successfully")
		}
	}()

	apphandler := &AppHandler{db: db}
	Router := mux.NewRouter().StrictSlash(true)
	Router.HandleFunc("/test/", apphandler.BasePageResponse).Methods("POST")

	// Using log.Print instead of log.Fatal allows defer to run
	if err := http.ListenAndServe(":8080", Router); err != nil {
		log.Printf("Server error: %v", err)
	}
}
