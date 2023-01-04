package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type InMemoryStore interface {
	SetIfNotExist(string, interface{}, time.Duration) error
	GetByKey(string) (string, error)
}

type LinkGenerator interface {
	Generate() string
}

type ShortenerHandler struct {
	store     InMemoryStore
	generator LinkGenerator
}

func NewShortenerHandler(store InMemoryStore, generator LinkGenerator) ShortenerHandler {
	return ShortenerHandler{
		store:     store,
		generator: generator,
	}
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

	r := s.generator.Generate()
	err = s.store.SetIfNotExist(r, link, time.Minute)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, r)
}

func (s ShortenerHandler) GetShortLink(c echo.Context) error {
	link := c.QueryParam("link")
	val, err := s.store.GetByKey(link)
	if err != nil {
		return err
	}

	if val == "" {
		c.NoContent(http.StatusNotFound)
	}

	return c.Redirect(http.StatusPermanentRedirect, val)
}
