package main

import (
	"github.com/khv1one/go-url-shortener/internal/api"
	"github.com/khv1one/go-url-shortener/internal/clients"
)

func main() {
	redisClient := clients.NewRedisClient()
	server := api.NewServer(redisClient)

	server.Logger.Fatal(server.Start(":8888"))
}
