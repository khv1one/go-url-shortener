package usecases

import (
	"fmt"
	"math/rand"
)

type LinkGenerator struct{}

func NewLinkGenerator() LinkGenerator {
	return LinkGenerator{}
}

func (g LinkGenerator) Generate() string {
	return fmt.Sprintf("link_%d", rand.Intn(100))
}
