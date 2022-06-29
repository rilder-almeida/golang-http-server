package main

// import (
// 	"context"

// 	"github.com/arquivei/foundationkit/app"
// 	fklog "github.com/arquivei/foundationkit/log"
// 	"github.com/rs/zerolog/log"
// )

// var version = "development"

// const (
// 	shutdownPriorityHTTP = app.ShutdownPriority(100)
// )

// func main() {
// 	app.SetupConfig(&config)

// 	ctx := fklog.SetupLoggerWithContext(context.Background(), config.Log, version)
// 	log.Ctx(ctx).Info().Str("config", fklog.Flatten(config)).Msg("Configuration")

// 	app.DefaultGracePeriod = config.Shutdown.Graceperiod
// 	app.DefaultShutdownTimeout = config.Shutdown.Timeout

// 	if err := app.NewDefaultApp(ctx); err != nil {
// 		log.Ctx(ctx).Fatal().Err(err).Msg("Failed to create app")
// 	}

// 	httpServer := getHTTPServer()

// 	app.RunAndWait(
// 		func() error {
// 			return httpServer.ListenAndServe()
// 		},
// 	)
// }

// // TODO
// // SEPARAR APIGET E APIINSERT
// // SEPARAR MAIN
