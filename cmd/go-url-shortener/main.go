package main

import "github.com/khv1one/go-url-shortener/internal/api"

func main() {
	server := api.NewServer()

	server.Logger.Fatal(server.Start(":8888"))
}
