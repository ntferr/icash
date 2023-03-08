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
	service crud.Service[entities.Bank]
}

type Controller interface {
	FindAll(c *fiber.Ctx) error
	FindByID(c *fiber.Ctx) error
	New(c *fiber.Ctx) error
	Alter(c *fiber.Ctx) error
	Remove(c *fiber.Ctx) error
}

func NewController(service crud.Service[entities.Bank]) Controller {
	return &bank{
		service: service,
	}
}

func (b bank) FindAll(c *fiber.Ctx) error {
	banks, err := b.service.GetAll(entities.Bank{})
	if err != nil {
		log.Printf("failed to retrive banks: %s", err.Error())
		return err
	}

	return c.JSON(&banks)
}

func (b bank) FindByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := snowflake.Validate(id); err != nil {
		log.Printf("failed to validate id: %e", err)
		return http_err.BadRequest(err)
	}

	bank, err := b.service.Get(entities.Bank{}, id)
	if err != nil {
		log.Printf("failed to get bank by id %s: %e", id, err)
		return http_err.InternalServerError(err)
	}

	return c.JSON(&bank)
}

func (b bank) New(c *fiber.Ctx) error {
	var bank entities.Bank
	if err := c.App().Config().JSONDecoder(c.Body(), &bank); err != nil {
		log.Printf("failed to decode bank: %e", err)
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

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("bank %s succefuly added ", bank.Name),
	})
}

func (b bank) Alter(c *fiber.Ctx) error {
	var bank entities.Bank
	id := c.Params("id")
	if err := snowflake.Validate(id); err != nil {
		log.Printf("failed to validate id: %e", err)
		return http_err.BadRequest(err)
	}

	if err := c.App().Config().JSONDecoder(c.Body(), &bank); err != nil {
		log.Printf("failed to decode bank: %e", err)
		return http_err.BadRequest(err)
	}

	bank.ID = id

	err := b.service.Update(bank)
	if err != nil {
		log.Printf("failed to update database: %e", err)
		return err
	}

	return c.JSON(fiber.Map{"message": "bank succefuly updated"})
}

func (b bank) Remove(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := snowflake.Validate(id); err != nil {
		log.Printf("failed to validate id: %e", err)
		return http_err.BadRequest(err)
	}

	err := b.service.Delete(entities.Bank{ID: id})
	if err != nil {
		log.Printf("failed to delete %s: %e", id, err)
		return http_err.InternalServerError(err)
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("bank %s succefuly deleted", id),
	})
}
