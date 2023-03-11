package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if _, err := os.Stat("sotr.db"); os.IsExist(err) {
		fmt.Printf("db already exist\n")
	} else {
		file, err := os.Create("sotr.db")
		if err != nil {

			fmt.Printf("%v\n", err)
		} else {
			fmt.Printf("We create db\n")
		}
		file.Close()
	}

	db, err := sql.Open("sqlite3", "sotr.db")
	if err == nil {
		quer, err := db.Prepare("CREATE TABLE sotr (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 'name' VARCHAR(20))")
		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			quer.Exec()
		}
	} else {
		fmt.Printf("%v\n", err)
	}

}
