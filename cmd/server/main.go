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
	defer func() {
		if err := app.Close(); err != nil {
			log.Printf("close app failed: %v\n", err)
		}
	}()

	// TODO: graceful shutdown
	if err := app.Run(); err != nil {
		log.Printf("app downed: %v\n", err)
	}
}
