package infrastructure

import (
	"database/sql"
	"log"
)

type DB struct {
	DBName     string
	Connection *sql.DB
}

func NewDB() *DB {
	db := &DB{
		DBName: "./data/todos.db",
	}
	db.getDBConnection()
	return db
}

func (d *DB) getDBConnection() {
	db, err := sql.Open("sqlite3", d.DBName)
	if err != nil {
		log.Fatal(err)
	}
	d.Connection = db

	_, err = d.Connection.Exec(`CREATE TABLE IF NOT EXISTS todos (
		id integer primary key,
		title text not null,
		content text not null,
		date TIMESTAMP DEFAULT (datetime(CURRENT_TIMESTAMP,'localtime'))
	);`)
	if err != nil {
		log.Fatal(err)
	}
}
