package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func HelloWorld(c *fiber.Ctx) error {
	err := c.SendString("Hello, World!")
	if err != nil {
		return err
	}

	return nil
}
