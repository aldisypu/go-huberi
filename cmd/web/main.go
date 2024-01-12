package main

import (
	"fmt"

	"github.com/aldisypu/go-huberi/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	config.NewDatabase(viperConfig, log)
	config.NewValidator(viperConfig)
	app := config.NewFiber(viperConfig)

	webPort := viperConfig.GetInt("web.port")
	err := app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
