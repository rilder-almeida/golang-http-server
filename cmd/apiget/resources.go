package main

import (
	"errors"
	"net/http"

	"github.com/arquivei/foundationkit/app"
	"github.com/go-kit/kit/endpoint"
	"github.com/golang-http-server/entities/nfe"
	"github.com/golang-http-server/entities/nfe/impltnfe"
	"github.com/golang-http-server/services/get"
	"github.com/golang-http-server/services/get/apiget"
	"github.com/golang-http-server/services/get/implget"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	getEndpoint   endpoint.Endpoint
	nfeRepository nfe.Repository
	httpServer    *http.Server
)

func getHTTPServer() *http.Server {
	if httpServer == nil {
		router := mux.NewRouter()

		router.PathPrefix("/nfe").Handler(apiget.MakeHTTPHandler(getGetEndpoint()))
		router.PathPrefix("/metrics").Handler(promhttp.Handler())

		httpAddr := ":" + config.HTTP.Port
		httpServer = &http.Server{
			Addr:    httpAddr,
			Handler: router,
		}

		app.RegisterShutdownHandler(
			&app.ShutdownHandler{
				Name:     "http_server",
				Priority: shutdownPriorityHTTP,
				Handler:  httpServer.Shutdown,
				Policy:   app.ErrorPolicyAbort,
			})
	}

	if httpServer == nil {
		panic("http server failed to initialize")
	}

	return httpServer
}

func getGetEndpoint() endpoint.Endpoint {
	// IMPLEMENT endpoint.Chain()() middlewares
	if getEndpoint == nil {
		getEndpoint = apiget.MakeAPIGetEndpoint(get.NewService(implget.NewAdapter(NewNFeRepository())))
	}
	return getEndpoint
}

func GetConnection() *gorm.DB {
	dsn := "host=" + config.Postgresql.Host + " port=" + config.Postgresql.Port + " user=" + config.Postgresql.User + " password=" + config.Postgresql.Password + " dbname=" + config.Postgresql.Dbname + " sslmode=" + config.Postgresql.Sslmode
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func NewNFeRepository() nfe.Repository {
	if nfeRepository == nil {
		switch config.Repository.Type {
		case "FILE":
			nfeRepository = impltnfe.WrapRepositoryWithCache(impltnfe.NewNFeFileRepository(config.Repository.FilePath))
		case "MEMORY":
			nfeRepository = impltnfe.WrapRepositoryWithCache(impltnfe.NewNFeMemoryRepository())
		case "POSTGRESQL":
			nfeRepository = impltnfe.WrapRepositoryWithCache(impltnfe.NewNFePostgresqlRepository(GetConnection()))
		default:
			panic(errors.New("bad repository, check env variables"))
		}
	}
	return nfeRepository
}
