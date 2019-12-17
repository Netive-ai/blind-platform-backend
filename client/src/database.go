package main

func createTableAccount() {
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

	_ ,err := gHandlers.DB.Exec(sqlStatement)
	fatal(err, 1)
}

func deleteTable(tableName string) {
	sqlStatement := "DROP TABLE " + tableName

	_ ,err := gHandlers.DB.Exec(sqlStatement)
	fatal(err, 1)
}

func genericQuery(sqlStatement string) {
	_ ,err := gHandlers.DB.Exec(sqlStatement)
	fatal(err, 0)
}

func createUser(username string, password string, sex string, age string, uuid string) {

}

func deleteUser() {

}

func searchUser() {

}