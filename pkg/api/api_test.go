package api

import (
	"github.com/blind-platform/pkg/api/auth"
	"github.com/blind-platform/pkg/conf"
	typ "github.com/blind-platform/pkg/type"
	"github.com/blind-platform/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func myInit() (typ.Handlers) {
	var handlers typ.Handlers
	conf := conf.Load()
	// INIT Router
	handlers.Router = mux.NewRouter()
	// INIT Keys
	signBytes, err := ioutil.ReadFile(conf.RSA.PathPrivate)
	util.Fatal(err)
	handlers.RSA.Private, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	util.Fatal(err)
	verifyBytes, err := ioutil.ReadFile(conf.RSA.PathPublic)
	util.Fatal(err)
	handlers.RSA.Public, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	util.Fatal(err)
	return handlers
}

func Router() *mux.Router {
	handler := myInit()
	handler.Router.HandleFunc("/api/auth/signup", auth.SignUp).Methods("POST")
	handler.Router.HandleFunc("/api/auth/signin", func(w http.ResponseWriter, r *http.Request) {
		auth.SignIn(w, r, handler.RSA) }).Methods("POST")
	handler.Router.HandleFunc("/api/auth/restrict", func(w http.ResponseWriter, r *http.Request) {
		auth.RestrictedHandler(w, r, handler.RSA) }).Methods("POST")
	return handler.Router
}

func TestValidSignupEndpoint(t *testing.T) {
	rec := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/auth/signup?username=username&password=password", nil)
	Router().ServeHTTP(rec, request)
	if rec.Result().StatusCode != http.StatusOK {
		t.Errorf("expected status OK, got %v", rec.Result().Status)
	} else {
		log.Info("Test passed (good username & password)")
	}
}

func TestWrongSignupEndpoint(t *testing.T) {
	rec := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/auth/signup?username=username&password=", nil)
	Router().ServeHTTP(rec, request)
	if rec.Result().StatusCode != http.StatusBadRequest {
		t.Errorf("expected status BadRequest, got %v", rec.Result().Status)
	} else {
		log.Info("Test 1 passed (no password)")
	}

	rec = httptest.NewRecorder()
	request, _ = http.NewRequest("POST", "/api/auth/signup?username=&password=password", nil)
	Router().ServeHTTP(rec, request)
	if rec.Result().StatusCode != http.StatusBadRequest {
		t.Errorf("expected status BadRequest, got %v", rec.Result().Status)
	} else {
		log.Info("Test 2 passed (no username)")
	}

	rec = httptest.NewRecorder()
	request, _ = http.NewRequest("POST", "/api/auth/signup?username=&password=", nil)
	Router().ServeHTTP(rec, request)
	if rec.Result().StatusCode != http.StatusBadRequest {
		t.Errorf("expected status BadRequest, got %v", rec.Result().Status)
	} else {
		log.Info("Test 3 passed (no username & password)")
	}
}

func TestValidSignIn(t *testing.T) {
	rec := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/auth/signin?username=test&password=known", nil)
	Router().ServeHTTP(rec, request)
	if rec.Result().StatusCode != http.StatusOK {
		t.Errorf("expected status OK, got %v", rec.Result().Status)
	} else {
		log.Info("Test passed (good username & password)")
	}
}

func TestWrongSignIn(t *testing.T) {
	rec := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/auth/signin?username=test&password=bad", nil)
	Router().ServeHTTP(rec, request)
	if rec.Result().StatusCode != http.StatusForbidden {
		t.Errorf("expected status BadRequest, got %v", rec.Result().Status)
	} else {
		log.Info("Test 1 passed (bad password)")
	}

	rec = httptest.NewRecorder()
	request, _ = http.NewRequest("POST", "/api/auth/signin?username=bad&password=known", nil)
	Router().ServeHTTP(rec, request)
	if rec.Result().StatusCode != http.StatusForbidden {
		t.Errorf("expected status BadRequest, got %v", rec.Result().Status)
	} else {
		log.Info("Test 2 passed (bad username)")
	}

	rec = httptest.NewRecorder()
	request, _ = http.NewRequest("POST", "/api/auth/signin?username=bad&password=bad", nil)
	Router().ServeHTTP(rec, request)
	if rec.Result().StatusCode != http.StatusForbidden {
		t.Errorf("expected status BadRequest, got %v", rec.Result().Status)
	} else {
		log.Info("Test 3 passed (bad username & password)")
	}
}