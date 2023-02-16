package bank

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ntferr/icash/entities"
	"github.com/ntferr/icash/http"
	"github.com/ntferr/icash/pkg/snowflake"
	banks "github.com/ntferr/icash/service/bank"
)

type bank struct {
	service banks.Service
}

type Controller interface {
	FindBankByID(c *fiber.Ctx) error
	NewBank(c *fiber.Ctx) error
}

func NewController(service banks.Service) Controller {
	return &bank{
		service: service,
	}
}

func (b bank) FindBankByID(c *fiber.Ctx) error {
	id := c.Params("bank_id")
	err := snowflake.Validate(id)
	if err != nil {
		log.Printf("failed to validate id: %e", err)
		return http.BadRequest(err)
	}

	bank, err := b.service.Get(id)
	if err != nil {
		log.Printf("failed to get bank by id %s: %e", id, err)
		return http.InternalServerError(err)
	}

	value, err := c.App().Config().JSONEncoder(bank)
	if err != nil {
		log.Printf("failed to encode bank into json: %e", err)
		return http.InternalServerError(err)
	}

	return c.JSON(value)
}

func (b bank) NewBank(c *fiber.Ctx) error {
	var bank entities.Bank
	value := c.Body()
	if err := c.App().Config().JSONDecoder(value, bank); err != nil {
		log.Printf("failed to decode bank: %e", err)
		return http.BadRequest(err)
	}

	id, err := snowflake.GenerateNew()
	if err != nil {
		log.Printf("failed to generate id: %f", err)
		return http.InternalServerError(err)
	}
	bank.ID = *id

	err = b.service.Insert(&bank)
	if err != nil {
		return http.InternalServerError(err)
	}

	return nil
}
