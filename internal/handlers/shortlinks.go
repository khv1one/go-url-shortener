package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateShortLink(c echo.Context) error {
	return c.String(http.StatusOK, "Post\n")
}

func GetShortLink(c echo.Context) error {
	return c.String(http.StatusOK, "Get\n")
}
