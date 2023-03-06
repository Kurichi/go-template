package main

import (
	"app/api/router"
	"context"
	"os"
	"os/signal"
	"time"
)

func main() {
	e := router.NewRouter()

	e.Logger.Fatal(e.Start(":8080"))

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
