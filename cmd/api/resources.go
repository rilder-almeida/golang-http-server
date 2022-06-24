package main

import (
	"errors"
	"net/http"

	"github.com/arquivei/foundationkit/app"
	"github.com/golang-http-server/entities/nfe"
	"github.com/golang-http-server/entities/nfe/impltnfe"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getHTTPServer() *http.Server {
	r := mux.NewRouter()

	handler := WrapHandlerWithMetrics(NewHandler())
	r.PathPrefix("/nfe/v1").Handler(handler)
	r.PathPrefix("/metrics").Handler(promhttp.Handler())

	httpAddr := ":" + config.HTTP.Port
	httpServer := &http.Server{
		Addr:    httpAddr,
		Handler: r,
	}

	app.RegisterShutdownHandler(
		&app.ShutdownHandler{
			Name:     "http_server",
			Priority: shutdownPriorityHTTP,
			Handler:  httpServer.Shutdown,
			Policy:   app.ErrorPolicyAbort,
		})

	return httpServer
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
	switch config.Repository.Type {
	case "INFILE":
		return impltnfe.WrapRepositoryWithCache(impltnfe.NewNFeInfileRepository(config.Repository.FilePath))
	case "INMEMORY":
		return impltnfe.WrapRepositoryWithCache(impltnfe.NewNFeInMemoryRepository())
	case "POSTGRESQL":
		return impltnfe.WrapRepositoryWithCache(impltnfe.NewNFePostgresqlRepository(GetConnection()))
	default:
		panic(errors.New("bad repository, check env variables"))
	}
}
