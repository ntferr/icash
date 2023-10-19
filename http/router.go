package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ntferr/icash/controller"
	"github.com/ntferr/icash/controller/bank"
	"github.com/ntferr/icash/controller/card"
	"github.com/ntferr/icash/controller/debt"
	"github.com/ntferr/icash/controller/ticket"
)

func SetupRouter(api fiber.Router, c controller.Controllers) {
	api.Get("/health-check", c.Health.Status)

	bankRouter(api.Group("/banks"), c.Bank)
	cardRouter(api.Group("/cards"), c.Card)
	ticketRouter(api.Group("/tickets"), c.Ticket)
	debtRouter(api.Group("/debts"), c.Debt)
}

func bankRouter(router fiber.Router, b bank.Controller) {
	router.Get("/", b.FindAll)
	router.Get("/:id", b.FindByID)
	router.Post("", b.New)
	router.Put("/:id", b.Alter)
	router.Delete("/:id", b.Remove)
}

func cardRouter(router fiber.Router, c card.Controller) {
	router.Get("/", c.FindAll)
	router.Get("/:id", c.FindByID)
	router.Post("", c.New)
	router.Put("/:id", c.Alter)
	router.Delete("/:id", c.Remove)
}

func ticketRouter(router fiber.Router, t ticket.Controller) {
	router.Get("/", t.FindAll)
	router.Get("/:id", t.FindByID)
	router.Post("", t.New)
	router.Put("/:id", t.Alter)
	router.Delete("/:id", t.Remove)
}

func debtRouter(router fiber.Router, d debt.Controller) {
	router.Get("/", d.FindAll)
	router.Get("/:id", d.FindByID)
	router.Post("", d.New)
	router.Put("/:id", d.Alter)
	router.Delete("/:id", d.Remove)
}
