package ticket

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ntferr/icash/entities"
	http_err "github.com/ntferr/icash/http/error"
	"github.com/ntferr/icash/pkg/snowflake"
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

func (t ticket) FindAll(ctx *fiber.Ctx) error {
	tickets, err := t.service.FindAll(entities.Ticket{})
	if err != nil {
		log.Printf("failed to find all tickets %e", err)
		return http_err.NotFound(err)
	}

	return ctx.JSON(&tickets)
}

func (t ticket) FindByID(ctx *fiber.Ctx) error {
	ticketID := ctx.Params("id")
	if err := snowflake.Validate(ticketID); err != nil {
		log.Printf("failed to validate ticketID")
		return http_err.BadRequest(err)
	}

	ticket, err := t.service.FindByID(entities.Ticket{ID: ticketID})
	if err != nil {
		log.Printf("failed to find ticket by id %s : %e", ticketID, err)
		return http_err.NotFound(err)
	}

	return ctx.JSON(&ticket)
}

func (t ticket) New(ctx *fiber.Ctx) error {
	var ticket entities.Ticket
	if err := ctx.App().Config().JSONDecoder(ctx.Body(), ticket); err != nil {
		log.Printf("failed to decode ticket: %e", err)
		return http_err.BadRequest(err)
	}

	ticketID, err := snowflake.GenerateNew()
	if err != nil {
		log.Printf("failed to generate id %f", err)
		return http_err.InternalServerError(err)
	}

	ticket.ID = *ticketID
	if err := t.service.Insert(ticket); err != nil {
		log.Printf("failed to insert ticket %s: %e", *ticketID, err)
		return http_err.InternalServerError(err)
	}

	return ctx.JSON(fiber.Map{
		"message": fmt.Sprintf("ticket %s succefuly added", *ticketID),
	})
}

func (t ticket) Alter(ctx *fiber.Ctx) error {
	var ticket entities.Ticket
	ticketID := ctx.Params("id")
	if err := snowflake.Validate(ticketID); err != nil {
		log.Printf("failed to validate ticket %s", ticketID)
		return http_err.BadRequest(err)
	}

	if err := ctx.App().Config().JSONDecoder(ctx.Body(), ticket); err != nil {
		log.Printf("failed to decode ticket %e", err)
		return http_err.BadRequest(err)
	}

	ticket.ID = ticketID
	if err := t.service.Update(ticket); err != nil {
		log.Printf("failed to update tikcet %s", ticketID)
		return http_err.InternalServerError(err)
	}

	return ctx.JSON(fiber.Map{
		"message": fmt.Sprintf("ticket %s succefuly updated", ticketID),
	})

}

func (t ticket) Remove(ctx *fiber.Ctx) error {
	ticketID := ctx.Params("id")
	if err := snowflake.Validate(ticketID); err != nil {
		log.Printf("failed to validate id: %e", err)
		return http_err.BadRequest(err)
	}

	if err := t.service.Delete(entities.Ticket{ID: ticketID}); err != nil {
		log.Printf("failed to delete ticket %s", ticketID)
		return http_err.InternalServerError(err)
	}

	return ctx.JSON(fiber.Map{
		"message": fmt.Sprintf("ticket %s succefuly deleted", ticketID),
	})
}
