package main

import (
	"github.com/Kurichi/go-template/pkg/config"
	"github.com/Kurichi/go-template/provider"
)

func main() {
	cfg := config.New()

	e := provider.InitServer()
	e.Logger.Fatal(e.Start(":" + cfg.ServerPort))
}
