package main

import (
	"github.com/tora0091/react-todo/server/infrastructure"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := infrastructure.NewDB()

	r := infrastructure.NewRouting(db.Connection)
	r.Run()

	defer db.Connection.Close()
}
