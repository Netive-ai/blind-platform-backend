package main

// Example HTTP auth using asymmetric crypto/RSA keys
// This is based on a (now outdated) example at https://gist.github.com/cryptix/45c33ecf0ae54828e63b

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"io/ioutil"
)


var (
	gHandlers Handlers
	users = map[string]string{
		"test": "known",
	}
)

func Connect(conf Conf) {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		conf.AwsDB.Endpoint,
		5432,
		conf.AwsDB.User,
		conf.AwsDB.Password,
		conf.AwsDB.DBName)
	gHandlers.DB, err = sql.Open("postgres", psqlInfo)
	fatal(err)
	err = gHandlers.DB.Ping()
	fatal(err)
	createTableAccount()
}

// read the key files before starting http handlers
func init() {
	conf := loadConf()
	// INIT Router
	gHandlers.Router = mux.NewRouter()
	// INIT Keys
	signBytes, err := ioutil.ReadFile(conf.RSA.PathPrivate)
	fatal(err)
	gHandlers.RSA.Private, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)
	verifyBytes, err := ioutil.ReadFile(conf.RSA.PathPublic)
	fatal(err)
	gHandlers.RSA.Public, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)
	// INIT DB
	Connect(conf)
}

func main() {
	handleRequests()
	defer gHandlers.DB.Close()
}