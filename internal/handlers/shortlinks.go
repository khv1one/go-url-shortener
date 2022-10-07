package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/khv1one/go-url-shortener/internal/clients"
	"github.com/labstack/echo/v4"
)

type ShortenerHandler struct {
	RedisClient clients.RedisClient
}

func (s ShortenerHandler) CreateShortLink(c echo.Context) error {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return err
	}
	link, has := json_map["link"]
	if !has {
		return c.NoContent(http.StatusBadRequest)
	}

	r := fmt.Sprintf("link_%d", rand.Intn(100))
	_, err = s.RedisClient.SetNX(r, link, 0).Result()
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, r)
}

func (s ShortenerHandler) GetShortLink(c echo.Context) error {
	link := c.QueryParam("link")
	val, err := s.RedisClient.Get(link).Result()
	if err != nil {
		return err
	}

	if val == "" {
		c.NoContent(http.StatusNotFound)
	}

	return c.Redirect(http.StatusPermanentRedirect, val)
}
