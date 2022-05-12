package impltnfe

import (
	"fmt"

	"github.com/golang-http-server/entities/nfe"
)

type Config struct {
	RepositoryType     string
	RepositoryFilePath string
}

func NewNfeRepository(c Config) nfe.Repository {
	switch c.RepositoryType {
	case "INFILE":
		return NewNfeInfileRepository(c.RepositoryFilePath)
	case "INMEMORY":
		return NewNfeInMemoryRepository()
	default:
		panic(fmt.Sprintf("Bad repository: %s", c))
	}
}
