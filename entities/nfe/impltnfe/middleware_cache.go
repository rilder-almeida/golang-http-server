package impltnfe

import (
	"github.com/golang-http-server/entities/nfe"
	"github.com/golang-http-server/shared"
)

type cacheMiddleware struct {
	next  nfe.Repository
	cache shared.CacheSync
}

func WrapRepositoryWithCache(repository nfe.Repository) nfe.Repository {
	return &cacheMiddleware{
		next:  repository,
		cache: shared.NewCacheSync(),
	}
}

func (cp *cacheMiddleware) FindByID(id string) (nfe.NFeDocument, error) {
	if value, ok := cp.cache.Load(id); ok {
		return value.(nfe.NFeDocument), nfe.ErrAlreadyExists
	}
	nfeDocument, err := cp.next.FindByID(id)

	if err == nfe.ErrAlreadyExists {
		cp.cache.Store(id, nfeDocument)
	}
	return nfeDocument, err
}

func (cp *cacheMiddleware) Save(nfeDocument nfe.NFeDocument) error {
	return cp.next.Save(nfeDocument)
}
