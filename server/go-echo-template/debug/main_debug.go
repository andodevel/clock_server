package main

import (
	"github.com/andodevel/go-echo-template/bootstrap"
	"github.com/andodevel/go-echo-template/db"
	server "github.com/andodevel/go-echo-template/server/go-echo-template"
)

func main() {
	bootstrap.EnableDebug()
	bootstrap.Init()
	db.Init()
	server.Start()
}
