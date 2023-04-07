package ticket

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ntferr/icash/entities"
	"github.com/ntferr/icash/service/crud"
)

type ticket struct {
	service crud.Contract[entities.Ticket]
}

type Controller interface {
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	New(ctx *fiber.Ctx) error
	Alter(ctx *fiber.Ctx) error
	Remove(ctx *fiber.Ctx) error
}

func NewController(service crud.Contract[entities.Ticket]) Controller {
	return &ticket{
		service: service,
	}
}

func (t ticket) FindAll(ctx *fiber.Ctx) error

func (t ticket) FindByID(ctx *fiber.Ctx) error

func (t ticket) New(ctx *fiber.Ctx) error

func (t ticket) Alter(ctx *fiber.Ctx) error

func (t ticket) Remove(ctx *fiber.Ctx) error
