package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ntferr/icash/controller"
)

type ControllerInterface interface {
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	New(ctx *fiber.Ctx) error
	Alter(ctx *fiber.Ctx) error
	Remove(ctx *fiber.Ctx) error
}

func SetupRouter(api fiber.Router, c controller.Controllers) {
	api.Get("/health-check", c.Health.Status)

	resgisterRoutes(api.Group("/banks"), c.Bank)
	resgisterRoutes(api.Group("/cards"), c.Card)
	resgisterRoutes(api.Group("/tickets"), c.Ticket)
	resgisterRoutes(api.Group("/debts"), c.Debt)
}

func resgisterRoutes(router fiber.Router, ctrl ControllerInterface) {
	router.Get("/", ctrl.FindAll)
	router.Get("/:id", ctrl.FindByID)
	router.Post("", ctrl.New)
	router.Put("/:id", ctrl.Alter)
	router.Delete("/:id", ctrl.Remove)
}
