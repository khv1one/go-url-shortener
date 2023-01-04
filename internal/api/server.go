package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ShortenerHandler interface {
	GetShortLink(echo.Context) error
	CreateShortLink(echo.Context) error
}

type Router struct {
	*echo.Echo
}

func NewRouter(shortenerHandler ShortenerHandler) Router {
	router := Router{echo.New()}

	// Middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	router.RoutingShortenerApi(shortenerHandler)

	return router
}

func (s Router) RoutingShortenerApi(handler ShortenerHandler) {
	s.GET("/link", handler.GetShortLink)
	s.POST("/link", handler.CreateShortLink)
}
