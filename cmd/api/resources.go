package main

// import (
// 	"errors"
// 	"net/http"

// 	"github.com/arquivei/foundationkit/app"
// 	"github.com/golang-http-server/entities/nfe"
// 	"github.com/golang-http-server/entities/nfe/impltnfe"
// 	"github.com/gorilla/mux"
// 	"github.com/prometheus/client_golang/prometheus/promhttp"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// func getHTTPServer() *http.Server {
// 	r := mux.NewRouter()

// 	handler := WrapHandlerWithMetrics(NewHandler())
// 	r.PathPrefix("/nfe/get").Handler(handler)
// 	// TODO
// 	// r.PathPrefix("/nfe/get").Handler(getGetEndpoint())
// 	// TODO
// 	// r.PathPrefix("/nfe/insert").Handler(getInsertEndpoint())
// 	r.PathPrefix("/metrics").Handler(promhttp.Handler())

// 	httpAddr := ":" + config.HTTP.Port
// 	httpServer := &http.Server{
// 		Addr:    httpAddr,
// 		Handler: r,
// 	}

// 	app.RegisterShutdownHandler(
// 		&app.ShutdownHandler{
// 			Name:     "http_server",
// 			Priority: shutdownPriorityHTTP,
// 			Handler:  httpServer.Shutdown,
// 			Policy:   app.ErrorPolicyAbort,
// 		})

// 	return httpServer
// }

// // IMPLEMENT
// // func getGetEndpoint() endpoint.Endpoint {}
// // func getInsertEndpoint() endpoint.Endpoint {}

// func GetConnection() *gorm.DB {
// 	dsn := "host=" + config.Postgresql.Host + " port=" + config.Postgresql.Port + " user=" + config.Postgresql.User + " password=" + config.Postgresql.Password + " dbname=" + config.Postgresql.Dbname + " sslmode=" + config.Postgresql.Sslmode
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	return db
// }

// func NewNFeRepository() nfe.Repository {
// 	switch config.Repository.Type {
// 	case "FILE":
// 		return impltnfe.WrapRepositoryWithCache(impltnfe.NewNFeFileRepository(config.Repository.FilePath))
// 	case "MEMORY":
// 		return impltnfe.WrapRepositoryWithCache(impltnfe.NewNFeMemoryRepository())
// 	case "POSTGRESQL":
// 		return impltnfe.WrapRepositoryWithCache(impltnfe.NewNFePostgresqlRepository(GetConnection()))
// 	default:
// 		panic(errors.New("bad repository, check env variables"))
// 	}
// }
