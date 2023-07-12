package bank

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ntferr/icash/entities"
	http_err "github.com/ntferr/icash/http/error"
	"github.com/ntferr/icash/pkg/snowflake"
	"github.com/ntferr/icash/service/crud"
)

type bank struct {
	service crud.Contract[entities.Bank]
}

type Controller interface {
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	New(ctx *fiber.Ctx) error
	Alter(ctx *fiber.Ctx) error
	Remove(ctx *fiber.Ctx) error
}

func NewController(service crud.Contract[entities.Bank]) Controller {
	return &bank{
		service: service,
	}
}

func (b bank) FindAll(ctx *fiber.Ctx) error {
	banks, err := b.service.FindAll(entities.Bank{})
	if err != nil {
		log.Printf("failed to retrive banks: %s", err.Error())
		return http_err.NotFound(err)
	}

	return ctx.JSON(&banks)
}

func (b bank) FindByID(ctx *fiber.Ctx) error {
	bankId := ctx.Params("id")
	if err := snowflake.Validate(bankId); err != nil {
		log.Printf("failed to validate id: %e", err)
		return http_err.BadRequest(err)
	}

	bank, err := b.service.FindByID(entities.Bank{ID: bankId})
	if err != nil {
		log.Printf("failed to get bank by id %s: %e", bankId, err)
		return http_err.NotFound(err)
	}

	return ctx.JSON(&bank)
}

func (b bank) New(ctx *fiber.Ctx) error {
	var bank entities.Bank
	if err := ctx.App().Config().JSONDecoder(ctx.Body(), &bank); err != nil {
		log.Printf("failed to decode bank: %e", err)
		return http_err.BadRequest(err)
	}

	if err := bank.Validate(); err != nil {
		log.Printf("validation has failed: %e", err)
		return http_err.BadRequest(err)
	}

	id, err := snowflake.GenerateNew()
	if err != nil {
		log.Printf("failed to generate id: %f", err)
		return http_err.InternalServerError(err)
	}
	bank.ID = *id

	err = b.service.Insert(bank)
	if err != nil {
		log.Printf("failed to insert bank: %f", err)
		return http_err.InternalServerError(err)
	}

	return ctx.JSON(fiber.Map{
		"message": fmt.Sprintf("bank %s succefuly added ", bank.Name),
	})
}

func (b bank) Alter(ctx *fiber.Ctx) error {
	var bank entities.Bank
	bankId := ctx.Params("id")
	if err := snowflake.Validate(bankId); err != nil {
		log.Printf("failed to validate id: %e", err)
		return http_err.BadRequest(err)
	}

	if err := ctx.App().Config().JSONDecoder(ctx.Body(), &bank); err != nil {
		log.Printf("failed to decode bank: %e", err)
		return http_err.BadRequest(err)
	}

	bank.ID = bankId
	if err := b.service.Update(bank); err != nil {
		log.Printf("failed to update bank: %e", err)
		return http_err.InternalServerError(err)
	}

	return ctx.JSON(fiber.Map{
		"message": "bank succefuly updated",
	})
}

func (b bank) Remove(ctx *fiber.Ctx) error {
	bankId := ctx.Params("id")
	if err := snowflake.Validate(bankId); err != nil {
		log.Printf("failed to validate id: %e", err)
		return http_err.BadRequest(err)
	}

	err := b.service.Delete(entities.Bank{ID: bankId})
	if err != nil {
		log.Printf("failed to delete bank %s: %e", bankId, err)
		return http_err.InternalServerError(err)
	}

	return ctx.JSON(fiber.Map{
		"message": fmt.Sprintf("bank %s succefuly deleted", bankId),
	})
}
