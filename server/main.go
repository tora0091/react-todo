package main

import (
	"github.com/tora0091/react-todo/server/infrastructure"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := infrastructure.NewDB()
	route := infrastructure.NewRouting(db.Connection)
	route.Run()

	defer db.Connection.Close()
}
