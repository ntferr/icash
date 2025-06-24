package error

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const (
	createSnowflakeResponse = "An error ocurred to create an snowflake id"
	notFindAllBanksResponse = "An error ocurred to get all banks"
	insertBankResponse      = "An error ocurred to insert bank"
	defaultResponse         = "An error ocurred into the API"
)

func WriteResponseError(err error) error {
	if errors.Is(err, project_errors.ErrToFind) {
		return fiber.NewError(http.StatusInternalServerError, notFindAllBanksResponse, err.Error())
	} else if errors.Is(err, project_errors.ErrToCreateSnowflake) {
		return fiber.NewError(http.StatusInternalServerError, createSnowflakeResponse, err.Error())
	} else if errors.Is(err, project_errors.ErrInsertBank) {
		return fiber.NewError(http.StatusInternalServerError, insertBankResponse, err.Error())
	} else {
		return fiber.NewError(http.StatusInternalServerError, defaultResponse)
	}
}
