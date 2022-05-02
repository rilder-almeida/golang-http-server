package impltnfe

import (
	"fmt"

	"github.com/golang-http-server/entities/nfe"
)

type Config struct {
	Repository   string
	JsonFilePath string
}

func NewNfeRepository(c Config) nfe.Repository {
	switch c.Repository {
	case "INFILE":
		return NewNfeInfileRepository(c.JsonFilePath)
	case "INMEMORY":
		return NewNfeInMemoryRepository()
	default:
		panic(fmt.Sprintf("Bad repository: %s", c))
	}
}
