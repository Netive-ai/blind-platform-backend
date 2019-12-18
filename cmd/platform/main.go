package main

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jackline/pkg/conf"
	"github.com/jackline/pkg/database"
	jlh "github.com/jackline/pkg/jl_http"
	"github.com/jackline/pkg/type"
	"github.com/jackline/pkg/util"
	_ "github.com/lib/pq"
	"io/ioutil"
)

func Connect(handlers typ.Handlers, conf conf.Conf) {

	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		conf.AwsDB.Endpoint,
		5432,
		conf.AwsDB.User,
		conf.AwsDB.Password,
		conf.AwsDB.DBName)
	handlers.DB, err = sql.Open("postgres", psqlInfo)
	util.Fatal(err)
	err = handlers.DB.Ping()
	util.Fatal(err)
	database.CreateTableAccount(handlers.DB)
}

// read the key files before starting http handlers
func my_init() (typ.Handlers) {
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
	// INIT DB
	Connect(handlers, conf)
	return handlers
}

func main() {
	handlers := my_init()

	jlh.HandleRequests(&handlers)
	defer handlers.DB.Close()
}