package sets

import (
	"github.com/google/wire"
	"github.com/khv1one/go-url-shortener/internal/api"
	"github.com/khv1one/go-url-shortener/internal/clients"
	"github.com/khv1one/go-url-shortener/internal/handlers"
	"github.com/khv1one/go-url-shortener/internal/usecases"
)

var (
	ClientsSet = wire.NewSet(
		clients.NewRedisClient,
		wire.Bind(new(handlers.InMemoryStore), new(clients.RedisClient)),
	)

	HandlersSet = wire.NewSet(
		handlers.NewShortenerHandler,
		wire.Bind(new(api.ShortenerHandler), new(handlers.ShortenerHandler)),
	)

	UsecasesSet = wire.NewSet(
		usecases.NewLinkGenerator,
		wire.Bind(new(handlers.LinkGenerator), new(usecases.LinkGenerator)),
	)
)
