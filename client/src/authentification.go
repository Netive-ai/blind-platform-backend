package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)


// Define some custom types were going to use within our tokens
type CustomerInfo struct {
	Name string
	Kind string
}

type CustomClaim struct {
	*jwt.StandardClaims
	TokenType string
	CustomerInfo
}

func signUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/jwt")
}

func createToken(user string) (string, error) {
	// create a signer for rsa 256
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	// set our claims
	t.Claims = &CustomClaim{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		},
		"level1",
		CustomerInfo{user, "human"},
	}
	return t.SignedString(gHandlers.RSA.Private)
}

// reads the form values, checks them and creates the token
func signIn(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	pass := r.FormValue("pass")

	log.Printf("Authenticate: user[%s] pass[%s]\n", user, pass)
	// check values
	if user != "test" || pass != "known" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Wrong info")
		return
	}
	tokenString, err := createToken(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Sorry, error while Signing Token!")
		log.Printf("Token Signing error: %v\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/jwt")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tokenString))
}

// only accessible with a valid token
func restrictedHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header["Bearer"] != nil {
		// Get token from request
		token, err := jwt.ParseWithClaims(r.Header["Bearer"][0], &CustomClaim{}, func(token *jwt.Token) (interface{}, error) {
			return gHandlers.RSA.Public, nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Invalid token:", err)
			return
		}
		// Token is valid
		fmt.Fprintln(w, "Welcome,", token.Claims.(*CustomClaim).Name)
		return
	} else {
		log.Info("No \"Bearer\" received in request on /api/auth/signup")
	}
}

