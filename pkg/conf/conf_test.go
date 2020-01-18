package conf

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestLoad(t *testing.T) {
	conf := Load()
	if len(conf.RSA.PathPrivate) == 0 {
		t.Error("no rsa private key loaded")
	}
	if len(conf.RSA.PathPublic) == 0 {
		t.Error("no rsa public key loaded")
	}
	if len(conf.AwsDB.User) == 0 {
		t.Error("no db user loaded")
	}
	if len(conf.AwsDB.Password) == 0 {
		t.Error("no db password loaded")
	}
	if len(conf.AwsDB.Endpoint) == 0 {
		t.Error("no db endpoint loaded")
	}
	if len(conf.AwsDB.DBName) == 0 {
		t.Error("no db name loaded")
	}
	log.Info("all the config informations got loaded")
}