package main

import (
	"superstore_api/routes"
)

func main()  {
	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}