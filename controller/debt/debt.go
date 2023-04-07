package debt

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ntferr/icash/entities"
	"github.com/ntferr/icash/service/crud"
)

type debt struct {
	service crud.Contract[entities.Debt]
}

type Controller interface {
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	New(ctx *fiber.Ctx) error
	Alter(ctx *fiber.Ctx) error
	Remove(ctx *fiber.Ctx) error
}

func NewController(service crud.Contract[entities.Debt]) Controller {
	return &debt{
		service: service,
	}
}

func (d debt) FindAll(ctx *fiber.Ctx) error

func (d debt) FindByID(ctx *fiber.Ctx) error

func (d debt) New(ctx *fiber.Ctx) error

func (d debt) Alter(ctx *fiber.Ctx) error

func (d debt) Remove(ctx *fiber.Ctx) error
