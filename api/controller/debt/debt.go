package debt

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	http_err "github.com/ntferr/icash/api/http/error"
	"github.com/ntferr/icash/entities"
	"github.com/ntferr/icash/pkg/snowflake"
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

func (d debt) FindAll(ctx *fiber.Ctx) error {
	debts, err := d.service.FindAll(entities.Debt{})
	if err != nil {
		log.Printf("failed to get all debts: %e", err)
		return http_err.NotFound(err)
	}

	return ctx.JSON(&debts)

}

func (d debt) FindByID(ctx *fiber.Ctx) error {
	debtID := ctx.Params("id")
	if err := snowflake.Validate(debtID); err != nil {
		log.Printf("failed to validate debt id: %e", err)
		return http_err.BadRequest(err)
	}

	debt, err := d.service.FindByID(entities.Debt{ID: debtID})
	if err != nil {
		log.Printf("failed to get debt by id %s: %e", debtID, err)
		return http_err.NotFound(err)
	}

	return ctx.JSON(&debt)
}

func (d debt) New(ctx *fiber.Ctx) error {
	var debt entities.Debt
	if err := ctx.App().Config().JSONDecoder(ctx.Body(), debt); err != nil {
		log.Printf("failed to decode debt: %e", err)
		return http_err.BadRequest(err)
	}

	id, err := snowflake.GenerateNew()
	if err != nil {
		log.Printf("failed to generate id %f", err)
		return http_err.InternalServerError(err)
	}

	debt.ID = *id
	if err = d.service.Insert(debt); err != nil {
		log.Printf("failed to insert debt: %e", err)
		return http_err.InternalServerError(err)
	}

	return ctx.JSON(fiber.Map{
		"message": fmt.Sprintf("card %s succefuly added", debt.ID),
	})
}

func (d debt) Alter(ctx *fiber.Ctx) error {
	var debt entities.Debt
	debtID := ctx.Params("id")
	if err := snowflake.Validate(debtID); err != nil {
		log.Printf("failed to validate debt id: %e", err)
		return http_err.BadRequest(err)
	}

	if err := ctx.App().Config().JSONDecoder(ctx.Body(), debt); err != nil {
		log.Printf("failed to decode debt: %e", err)
		return http_err.BadRequest(err)
	}

	debt.ID = debtID

	if err := d.service.Update(debt); err != nil {
		log.Printf("failed to update debt %s", debtID)
		return http_err.InternalServerError(err)
	}

	return ctx.JSON(fiber.Map{
		"message": "debt succefuly updated",
	})
}

func (d debt) Remove(ctx *fiber.Ctx) error {
	debtID := ctx.Params("id")
	if err := snowflake.Validate(debtID); err != nil {
		log.Printf("failed to validate debt id %e", err)
		return http_err.BadRequest(err)
	}

	if err := d.service.Delete(entities.Debt{ID: debtID}); err != nil {
		log.Printf("failed to delete debt %s", debtID)
		return http_err.InternalServerError(err)
	}

	return ctx.JSON(fiber.Map{
		"message": fmt.Sprintf("debt %s succefuly deleted", debtID),
	})
}
