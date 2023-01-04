//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/khv1one/go-url-shortener/internal/api"
	"github.com/khv1one/go-url-shortener/internal/pkg/wire/sets"
)

func InitializeApp() api.Router {
	wire.Build(
		sets.ClientsSet,
		sets.HandlersSet,
		sets.UsecasesSet,

		api.NewRouter,
	)

	return api.Router{}
}
