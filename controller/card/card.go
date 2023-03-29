package card

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ntferr/icash/entities"
	http_err "github.com/ntferr/icash/http/error"
	"github.com/ntferr/icash/pkg/snowflake"
	"github.com/ntferr/icash/service/crud"
)

type card struct {
	service crud.Contract[entities.Card]
}

type Controller interface {
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	New(ctx *fiber.Ctx) error
	Alter(ctx *fiber.Ctx) error
	Remove(ctx *fiber.Ctx) error
}

func NewController(service crud.Contract[entities.Card]) Controller {
	return &card{
		service: service,
	}
}

func (c card) FindAll(ctx *fiber.Ctx) error {
	cards, err := c.service.FindAll(entities.Card{})
	if err != nil {
		log.Printf("failed to get all cards: %e", err)
		return http_err.NotFound(err)
	}

	return ctx.JSON(&cards)
}

func (c card) FindByID(ctx *fiber.Ctx) error {
	cardId := ctx.Params("id")
	if err := snowflake.Validate(cardId); err != nil {
		log.Printf("failed to validate card id: %e", err)
		return http_err.BadRequest(err)
	}

	card, err := c.service.FindByID(entities.Card{ID: cardId})
	if err != nil {
		log.Printf("failed to get card by id %s: %e", cardId, err)
		return http_err.NotFound(err)
	}

	return ctx.JSON(&card)
}

func (c card) New(ctx *fiber.Ctx) error {
	var card entities.Card
	if err := ctx.App().Config().JSONDecoder(ctx.Body(), &card); err != nil {
		log.Printf("failed to decode card: %e", err)
		return http_err.BadRequest(err)
	}

	id, err := snowflake.GenerateNew()
	if err != nil {
		log.Printf("failed to generate id %f", err)
		return http_err.InternalServerError(err)
	}

	card.ID = *id

	if err = c.service.Insert(card); err != nil {
		log.Printf("failed to insert card: %e", err)
		return http_err.InternalServerError(err)
	}

	return ctx.JSON(fiber.Map{
		"message": fmt.Sprintf("card %s succefuly added", card.Number),
	})
}

func (c card) Alter(ctx *fiber.Ctx) error {
	var card entities.Card
	cardId := ctx.Params("id")
	if err := snowflake.Validate(cardId); err != nil {
		log.Printf("failed to validate id: %e", err)
		return http_err.BadRequest(err)
	}

	if err := ctx.App().Config().JSONDecoder(ctx.Body(), card); err != nil {
		log.Printf("failed to decode card: %e", err)
		return http_err.BadRequest(err)
	}

	card.ID = cardId
	if err := c.service.Update(card); err != nil {
		log.Printf("failed to update card: %e", err)
	}

	return ctx.JSON(fiber.Map{
		"message": "card succefuly updated",
	})
}

func (c card) Remove(ctx *fiber.Ctx) error {
	cardId := ctx.Params("id")
	if err := snowflake.Validate(cardId); err != nil {
		log.Printf("failed to validate id: %e", err)
		return http_err.BadRequest(err)
	}

	if err := c.service.Delete(entities.Card{ID: cardId}); err != nil {
		log.Printf("failed to delete card %s: %e", cardId, err)
		return http_err.InternalServerError(err)
	}

	return ctx.JSON(fiber.Map{
		"message": fmt.Sprintf("card %s succefuly deleted", cardId),
	})
}
