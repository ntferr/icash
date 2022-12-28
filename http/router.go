package http

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(api fiber.Router, c Controllers) {
	api.Get("/health-check", c.health.Status)
}
