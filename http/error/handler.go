package error

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ntferr/icash/project_errors"
)

const (
	notFindAllBanksResponse = "An error ocurred to get all banks"
	defaultResponse         = "An error ocurred into the API"
)

func WriteResponseError(err error) error {
	if errors.Is(err, project_errors.ErrToFind) {
		return fiber.NewError(http.StatusInternalServerError, notFindAllBanksResponse)
	} else {
		return fiber.NewError(http.StatusInternalServerError, defaultResponse)
	}
}
