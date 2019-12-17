package main

import (
	"log"
	"net/http"
)

func handleRequests() {
	gHandlers.Router.HandleFunc("/api/auth/signup", signUp).Methods("POST")
	gHandlers.Router.HandleFunc("/api/auth/signin", signIn).Methods("POST")
	gHandlers.Router.HandleFunc("/api/auth/restrict", restrictedHandler).Methods("POST")
	log.Fatal(http.ListenAndServe("localhost:8001", gHandlers.Router))
}