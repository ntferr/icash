package bank

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	http_err "github.com/icash/internal/http/error"
	"github.com/icash/internal/service"
	"github.com/icash/pkg/uuid"
)

type bank struct {
	service service.BankCrud
}

type Controller interface {
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	New(ctx *fiber.Ctx) error
	Alter(ctx *fiber.Ctx) error
	Remove(ctx *fiber.Ctx) error
}

func NewController(service service.BankCrud) Controller {
	return &bank{
		service: service,
	}
}

func (b bank) FindAll(ctx *fiber.Ctx) error {
	banks, err := b.service.FindAll()
	if err != nil {
		err = fmt.Errorf("%e: %w:",
			err,
			project_errors.ErrToFind,
		)
		log.Println(err)

		return http_err.WriteResponseError(err)
	}

	return ctx.JSON(&banks)
}

func (b bank) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := uuid.Validate(id); err != nil {
		err = fmt.Errorf("%e: %w",
			err,
			project_errors.ErrToValidateSnowflake,
		)
		log.Println(err)

		return http_err.WriteResponseError(err)
	}

	bank, err := b.service.FindByID(&id)
	if err != nil {
		err = fmt.Errorf("%e: %w",
			err,
			project_errors.ErrToFind,
		)
		log.Println(err)

		return http_err.WriteResponseError(err)
	}

	return ctx.JSON(&bank)
}

func (b bank) New(ctx *fiber.Ctx) error {
	var bank Bank
	if err := ctx.App().Config().JSONDecoder(ctx.Body(), &bank); err != nil {
		err = fmt.Errorf("%e: %w",
			err,
			project_errors.ErrToUnmarshal,
		)
		log.Println(err)

		return http_err.WriteResponseError(err)
	}

	if err := bank.Validate(); err != nil {
		err = fmt.Errorf("%e: %w",
			err,
			project_errors.ErrValidateBank,
		)
		log.Println(err)

		return http_err.WriteResponseError(err)
	}

	bank.ID = uuid.GenUUID5(&bank)

	if err := b.service.Insert(&bank); err != nil {
		err = fmt.Errorf("%e: %w",
			err,
			project_errors.ErrInsertBank,
		)
		log.Println(err)

		return http_err.WriteResponseError(err)
	}

	return ctx.JSON(fiber.Map{
		"message": fmt.Sprintf("bank %s succefuly added ", bank.Name),
	})
}

func (b bank) Alter(ctx *fiber.Ctx) error {
	var bank Bank
	id := ctx.Params("id")
	if err := uuid.Validate(id); err != nil {
		err = fmt.Errorf("%e: %w",
			err,
			project_errors.ErrToValidateSnowflake,
		)
		log.Println(err)

		return http_err.WriteResponseError(err)
	}

	if err := ctx.App().Config().JSONDecoder(ctx.Body(), &bank); err != nil {
		err = fmt.Errorf("%e: %w",
			err,
			project_errors.ErrToUnmarshal,
		)
		log.Println(err)

		return http_err.WriteResponseError(err)
	}

	bank.ID = id
	if err := b.service.Update(&bank); err != nil {
		err = fmt.Errorf("%e: %w",
			err,
			project_errors.ErrUpdateBank,
		)
		log.Println(err)

		return http_err.WriteResponseError(err)
	}

	return ctx.JSON(fiber.Map{
		"message": fmt.Sprintf("bank %s succefuly updated", id),
	})
}

func (b bank) Remove(ctx *fiber.Ctx) error {
	bankId := ctx.Params("id")
	if err := uuid.Validate(bankId); err != nil {
		log.Printf("failed to validate id: %e", err)
		return http_err.WriteResponseError(err)
	}

	err := b.service.Delete(&bankId)
	if err != nil {
		log.Printf("failed to delete bank %s: %e", bankId, err)
		return http_err.WriteResponseError(err)
	}

	return ctx.JSON(fiber.Map{
		"message": fmt.Sprintf("bank %s succefuly deleted", bankId),
	})
}
