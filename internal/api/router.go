package api

import "github.com/khv1one/go-url-shortener/internal/handlers"

func Routing(server Server) {
	server.GET("/link", handlers.GetShortLink)
	server.POST("/link", handlers.CreateShortLink)
}
