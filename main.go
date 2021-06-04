package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Dbconnection *sql.DB

func main() {
	// テーブル作成
	Dbconnection, _ := sql.Open("sqlite3", "example.sql")
	defer Dbconnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS person(
						name STRING,
						age INT)`
	_, err := Dbconnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// インサート
	cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	_, err = Dbconnection.Exec(cmd, "Mike", 24)
	if err != nil {
		log.Fatalln(err)
	}

	// アップデート
	cmd = "UPDATE person SET  age = ? WHERE name = ?"
	_, err = Dbconnection.Exec(cmd, 25, "Mike")
	if err != nil {
		log.Fatalln(err)
	}
}