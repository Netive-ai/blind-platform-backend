package jl_http

import (
	"github.com/blind-platform/pkg/auth"
	typ "github.com/blind-platform/pkg/type"
	"log"
	"net/http"
)

func HandleRequests(handler *typ.Handlers) {
	handler.Router.HandleFunc("/api/auth/signup", auth.SignUp).Methods("POST")
	handler.Router.HandleFunc("/api/auth/signin", func(w http.ResponseWriter, r *http.Request) {
		auth.SignIn(w, r, handler.RSA) }).Methods("POST")
	handler.Router.HandleFunc("/api/auth/restrict", func(w http.ResponseWriter, r *http.Request) {
		auth.RestrictedHandler(w, r, handler.RSA) }).Methods("POST")


	log.Fatal(http.ListenAndServe("localhost:8001", handler.Router))
}