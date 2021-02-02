package main

import (
	"ewarung-api-experiment/db"
	"ewarung-api-experiment/routes.go"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":9001"))
}
