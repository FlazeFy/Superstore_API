package main

import (
	"superstore_api/database"
	"superstore_api/routes"
)

func main() {
	database.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
