package main

import (
	"log"

	"github.com/Kurichi/go-template/internal/app"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Fatalf("initialize app failed: %v\n", err)
	}
	defer app.Close()

	// TODO: graceful shutdown
	if err := app.Run(); err != nil {
		log.Fatalf("app downed: %v\n", err)
	}
}
