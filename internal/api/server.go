package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	*echo.Echo
}

func NewServer() Server {
	server := Server{echo.New()}

	// Middleware
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	Routing(server)

	return server
}
