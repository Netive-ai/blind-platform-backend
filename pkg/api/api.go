package api

import (
	"fmt"
	"github.com/blind-platform/pkg/api/auth"
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

	fmt.Println("listen starting...")
	log.Fatal(http.ListenAndServe(":8001", handler.Router))
}