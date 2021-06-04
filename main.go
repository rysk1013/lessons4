package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Dbconnection *sql.DB

func main() {
	Dbconnection, _ := sql.Open("sqlite3", "example.sql")
	defer Dbconnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS person(
						name STRING,
						age INT)`
	_, err := Dbconnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}
}