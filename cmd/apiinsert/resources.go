package main

import (
	"errors"
	"net/http"

	"github.com/arquivei/foundationkit/app"
	"github.com/arquivei/foundationkit/contextmap"
	"github.com/arquivei/foundationkit/gokitmiddlewares/loggingmiddleware"
	"github.com/arquivei/foundationkit/gokitmiddlewares/metricsmiddleware"
	"github.com/arquivei/foundationkit/gokitmiddlewares/trackingmiddleware"
	"github.com/go-kit/kit/endpoint"
	"github.com/golang-http-server/entities/nfe"
	"github.com/golang-http-server/entities/nfe/impltnfe"
	"github.com/golang-http-server/services/insert"
	"github.com/golang-http-server/services/insert/apiinsert"
	"github.com/golang-http-server/services/insert/implinsert"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	insertEndpoint endpoint.Endpoint
	nfeRepository  nfe.Repository
	httpServer     *http.Server
)

func getHTTPServer() *http.Server {
	if httpServer == nil {
		router := mux.NewRouter()

		router.PathPrefix("/nfe").Handler(apiinsert.MakeHTTPHandler(getInsertEndpoint()))

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
		panic("http server is nil")
	}

	return httpServer
}

func getInsertEndpoint() endpoint.Endpoint {
	if insertEndpoint == nil {
		loggingConfig := loggingmiddleware.NewDefaultConfig("insert")
		metricsConfig := metricsmiddleware.NewDefaultConfig(instrumentingSystem, "insert")

		insertEndpoint = endpoint.Chain(
			contextmap.NewEndpointMiddleware(),
			trackingmiddleware.New(),
			loggingmiddleware.MustNew(loggingConfig),
			metricsmiddleware.MustNew(metricsConfig),
		)(apiinsert.MakeAPIInsertEndpoint(insert.NewService(implinsert.NewAdapter(NewNFeRepository()))))
	}
	return insertEndpoint
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
