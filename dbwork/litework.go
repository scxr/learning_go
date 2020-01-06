package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	}
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS users (id PRIMARY KEY, firstname TEXT, lastname TEXT)") // create the table
	statement.Exec()
	statement, _ = db.Prepare("INSERT INTO users (firstname, lastname) VALUES (?,?)") // prepare the statement to be executed
	statement.Exec("test", "user")                                                    // tells the statement what to put into the place holders
	ret, _ := db.Query("SELECT id, firstname, lastname FROM users")                   // select the statement

	// declare vars to be used by the fmt
	var id int
	var firstname string
	var lastname string

	for ret.Next() { // while theres values to be read, read them

		ret.Scan(&id, &firstname, &lastname)                                                      // scan the response for the values
		fmt.Println("ID : " + strconv.Itoa(id) + "firstname" + firstname + "lastname" + lastname) // print the vals
	}
}
