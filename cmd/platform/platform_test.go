package main_test

import (
	"github.com/blind-platform/cmd/platform"
	"github.com/blind-platform/pkg/api/auth"
	log "github.com/sirupsen/logrus"
	"testing"
)


func TestInit(t *testing.T) {
	handler := main.MyInit()
	result := handler.Router.HandleFunc("/api/auth/signup", auth.SignUp).Methods("POST").GetError()
	if result != nil {
		t.Error("Router not loaded")
	}
	log.Info("Router loaded")
	result = handler.DB.Ping()
	if result != nil {
		t.Error("DB not loaded")
	}
	log.Info("DB loaded")
	key := handler.RSA.Private
	if key == nil {
		t.Error("key not loaded")
	}
	log.Info("RSA loaded")
}