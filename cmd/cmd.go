package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"github.com/yourusername/fiber_home/internal/config"
	"github.com/yourusername/fiber_home/internal/http"
)

func main() {
	cfg := config.GetConfig()
	log.Infof("DB URL: %s", cfg.DBConfig.URL)

	app := fiber.New()
	http.ConfigureCommonMiddlewares(app)
	http.ConfigureRoutes(app)

	err := app.Listen(cfg.SrverPort.Port)
	if err != nil {
		log.Panicf("start server: %v", err)
	}
}
