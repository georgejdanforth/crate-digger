package main

import (
	"github.com/georgejdanforth/crate-digger/database"
	"github.com/georgejdanforth/crate-digger/server"
)

func main() {
	database.Init()
	server.Init()
}
