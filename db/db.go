package db

import (
	/*
		We need a special import syntax because we'll never directly use "github.com/mattn/go-sqlite3".
		Instead, it must be imported here so that GO uses it under the hood,
		because we will actually interact with the built-in SQL sub package "database/sql" thats part of the GO's standard library.
	*/
	"database/sql"

	/*
		This package will expose functionality thats used under the hood by the built-in SQL package "database/sql"
	*/
	_ "github.com/mattn/go-sqlite3" // '_' it must be part of the final project but we'll not use it directly, not letting it removed automatically.
)

var DB *sql.DB // Global variable, uppercase so that we can use it outside of this package.

/*
Func responsible for initializing the database
*/
func initDB() {
	/*
		Opening a connection: driver name, I named it "api.db" and it will be created automatically if it doesn't exist yet.
		All the data will be stored in "api.db" file and we will be able to interact with that data by using SQL queries.
	*/
	DB, err := sql.Open("sqlit3", "api.db")

	/*
		Checking Error
	*/
	if err != nil {
		panic("Could not connect to database.")
	}

	/*
		If no Error
	*/
	DB.SetMaxOpenConns(10) // Maximum connections can be opened simultaneously, lets set to 10
	DB.SetMaxIdleConns(5)  // Maximum connections opened if no ones using these connections at the moment, lets set to 5
}
