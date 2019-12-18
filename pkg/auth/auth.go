package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	typ "github.com/jackline/pkg/type"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// custom types were going to use within our tokens
type CustomerInfo struct {
	Name string
	Kind string
}

type CustomClaim struct {
	*jwt.StandardClaims
	TokenType string
	CustomerInfo
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/jwt")
}

func createToken(user string, rsa typ.RSA) (string, error) {
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
	return t.SignedString(rsa.Private)
}

// reads the form values, checks them and creates the token
func SignIn(w http.ResponseWriter, r *http.Request, rsa typ.RSA) {
	user := r.FormValue("user")
	pass := r.FormValue("pass")

	log.Printf("Authenticate: user[%s] pass[%s]\n", user, pass)
	// check values
	if user != "test" || pass != "known" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Wrong info")
		return
	}
	tokenString, err := createToken(user, rsa)
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
func RestrictedHandler(w http.ResponseWriter, r *http.Request, rsa typ.RSA) {
	if r.Header["Bearer"] != nil {
		// Get token from request
		token, err := jwt.ParseWithClaims(r.Header["Bearer"][0], &CustomClaim{}, func(token *jwt.Token) (interface{}, error) {
			return rsa.Public, nil
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

