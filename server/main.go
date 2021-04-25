package main

import (
	"database/sql"
	"log"

	"github.com/tora0091/react-todo/server/infrastructure"

	_ "github.com/mattn/go-sqlite3"
)

// create table todos (
// id integer primary key,
// title text not null,
// content text not null,
// date TIMESTAMP DEFAULT (datetime(CURRENT_TIMESTAMP,'localtime'))
// );
func getDBConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "./data/todos.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	dbConn := getDBConnection()

	r := infrastructure.NewRouting(dbConn)
	r.Run()

	defer dbConn.Close()
}
