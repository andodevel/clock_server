package main

import (
	"github.com/andodevel/clock_server/bootstrap"
	"github.com/andodevel/clock_server/db"
	server "github.com/andodevel/clock_server/server/clock_server"
)

func main() {
	bootstrap.EnableDebug()
	bootstrap.Init()
	db.Init()
	server.Start()
}
