package database

import (
	"database/sql"
	"github.com/jackline/pkg/util"
)

func CreateTableAccount(DB *sql.DB) {
	sqlStatement := `
	CREATE TABLE account(
	uuid serial PRIMARY KEY,
	username VARCHAR (50) UNIQUE NOT NULL,
	password VARCHAR (50) NOT NULL,
	email VARCHAR (355) UNIQUE NOT NULL,
	sex VARCHAR (1) NOT NULL,
	created_on TIMESTAMP NOT NULL,
	last_login TIMESTAMP
	);`

	_ ,err := DB.Exec(sqlStatement)
	util.Fatal(err, 1)
}

func DeleteTable(DB *sql.DB, tableName string) {
	sqlStatement := "DROP TABLE " + tableName

	_ ,err := DB.Exec(sqlStatement)
	util.Fatal(err, 1)
}

func GenericQuery(DB *sql.DB, sqlStatement string) {
	_ ,err := DB.Exec(sqlStatement)
	util.Fatal(err, 0)
}

func CreateUser(DB *sql.DB, username string, password string, sex string, age string, uuid string) {

}

func DeleteUser() {

}

func SearchUser() {

}