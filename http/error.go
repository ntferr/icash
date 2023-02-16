package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func InternalServerError(err error) error {
	return fiber.NewError(http.StatusInternalServerError, err.Error())
}

func NotFound(err error) error {
	return fiber.NewError(http.StatusNotFound, err.Error())
}

func BadRequest(err error) error {
	return fiber.NewError(http.StatusBadRequest, err.Error())
}
