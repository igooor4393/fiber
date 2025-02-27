package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/yourusername/fiber_home/internal/http/handlers"
)

func ConfigureCommonMiddlewares(app *fiber.App) {
	app.Use(recover.New())
}

func ConfigureRoutes(app *fiber.App) {
	apiGroup := app.Group("/api")
	apiGroup.Get("/world", handlers.HelloWorld)
}
