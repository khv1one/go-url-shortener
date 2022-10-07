package api

import (
	"github.com/khv1one/go-url-shortener/internal/clients"
	"github.com/khv1one/go-url-shortener/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	*echo.Echo
}

func NewServer(redisClient clients.RedisClient) Server {
	server := Server{echo.New()}

	// Middleware
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	server.RoutingShortenerApi(handlers.ShortenerHandler{RedisClient: redisClient})

	return server
}

func (s Server) RoutingShortenerApi(handler handlers.ShortenerHandler) {
	s.GET("/link", handler.GetShortLink)
	s.POST("/link", handler.CreateShortLink)
}
