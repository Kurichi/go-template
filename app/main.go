package main

import (
	"app/api/router"
)

func main() {
	e := router.NewRouter()

	e.Logger.Fatal(e.Start(":8080"))
}
