package typ

import (
	"crypto/rsa"
	"database/sql"
	"github.com/gorilla/mux"
)

/**
** OPERATION RELATED STRUCTURES
 */

type RSA struct {
	Private			*rsa.PrivateKey
	Public			*rsa.PublicKey
}

type Handlers struct {
	DB		*sql.DB
	Router	*mux.Router
	RSA 	RSA
}

/**
** USER RELATED STRUCTURES
 */
type User struct {
	id			string `json:"id"`
	firstName 	string `json:"firstName"`
	lastName	string `json:"lastName"`
	sex			string `json:"sex"`
	age			int    `json:"age"`
}