package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ntferr/icash/controller"
)

func SetupRouter(api fiber.Router, c controller.Controllers) {
	api.Get("/health-check", c.Health.Status)

	bankEndpoints(api.Group("/banks"), c)
	cardEndpoints(api.Group("/cards"), c)

}

func bankEndpoints(router fiber.Router, c controller.Controllers) {
	router.Get("/", c.Bank.FindAll)
	router.Get("/:id", c.Bank.FindByID)
	router.Post("", c.Bank.New)
	router.Put("/:id", c.Bank.Alter)
	router.Delete("/:id", c.Bank.Remove)
}

func cardEndpoints(router fiber.Router, c controller.Controllers) {
	router.Get("/", c.Card.FindAll)
	router.Get("/:id", c.Card.FindByID)
	router.Post("", c.Card.New)
	router.Put("/:id", c.Card.Alter)
	router.Delete("/:id", c.Card.Remove)
}
