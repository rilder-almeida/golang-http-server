package impltnfe

import (
	"errors"

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
	case "POSTGRESQL":
		return NewNfePostgresqlRepository()
	default:
		panic(errors.New("bad repository, check env variables"))
	}
}
