package application

import (
	"log"

	"anime/config"
	"anime/server"
	"anime/server/routes"
)

func Start(cfg *config.Config) {
	app := server.NewServer(cfg)

	routes.ConfigureRoutes(app)

	err := app.Start(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
