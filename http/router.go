package http

import "github.com/gofiber/fiber/v2"

func SetupRouter(api fiber.Router) {
	api.Get("/health-check", healthCheckHandler)
}
