package bank

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ntferr/icash/entities"
	http_err "github.com/ntferr/icash/http/error"
	"github.com/ntferr/icash/pkg/snowflake"
	banks "github.com/ntferr/icash/service/bank"
)

type bank struct {
	service banks.Service
}

type Controller interface {
	FindAll(c *fiber.Ctx) error
	FindByID(c *fiber.Ctx) error
	New(c *fiber.Ctx) error
}

func NewController(service banks.Service) Controller {
	return &bank{
		service: service,
	}
}

func (b bank) FindAll(c *fiber.Ctx) error {
	banks, err := b.service.GetAll()
	if err != nil {
		log.Printf("failed to retrive banks: %s", err.Error())
		return err
	}

	return c.JSON(&banks)
}

func (b bank) FindByID(c *fiber.Ctx) error {
	id := c.Params("id")
	err := snowflake.Validate(id)
	if err != nil {
		log.Printf("failed to validate id: %e", err)
		return http_err.BadRequest(err)
	}

	bank, err := b.service.Get(id)
	if err != nil {
		log.Printf("failed to get bank by id %s: %e", id, err)
		return http_err.InternalServerError(err)
	}

	return c.JSON(&bank)
}

func (b bank) New(c *fiber.Ctx) error {
	var bank entities.Bank
	value := c.Body()
	if err := c.App().Config().JSONDecoder(value, &bank); err != nil {
		log.Printf("failed to decode bank: %e", err)
		return http_err.BadRequest(err)
	}

	id, err := snowflake.GenerateNew()
	if err != nil {
		log.Printf("failed to generate id: %f", err)
		return http_err.InternalServerError(err)
	}
	bank.ID = *id

	err = b.service.Insert(&bank)
	if err != nil {
		return http_err.InternalServerError(err)
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("bank: %s succefuly added ", bank.Name),
	})
}
