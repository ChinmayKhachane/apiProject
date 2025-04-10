package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	db, err := CreateConn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	apphandler := &AppHandler{db: db}

	Router := mux.NewRouter().StrictSlash(true)
	Router.HandleFunc("/query/test", apphandler.FactDataTest).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", Router))

}
