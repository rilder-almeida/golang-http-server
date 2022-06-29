package main

import (
	"context"
	"runtime/debug"

	"github.com/arquivei/foundationkit/app"
	"github.com/arquivei/foundationkit/errors"
	fklog "github.com/arquivei/foundationkit/log"
	"github.com/arquivei/foundationkit/stringsutil"
	"github.com/rs/zerolog/log"
)

var version = "development"

const (
	shutdownPriorityHTTP = app.ShutdownPriority(100)
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal().
				Err(errors.NewFromRecover(r)).
				Str("panic_stack", stringsutil.Truncate(string(debug.Stack()), 1024)).
				Msg("App panicked")
		}
	}()

	app.SetupConfig(&config)

	ctx := fklog.SetupLoggerWithContext(context.Background(), config.Log, version)
	log.Ctx(ctx).Info().Str("config", fklog.Flatten(config)).Msg("Configuration")

	app.DefaultGracePeriod = config.Shutdown.Graceperiod
	app.DefaultShutdownTimeout = config.Shutdown.Timeout

	if err := app.NewDefaultApp(ctx); err != nil {
		log.Ctx(ctx).Fatal().Err(err).Msg("Failed to create app")
	}

	httpServer := getHTTPServer()

	app.RunAndWait(
		func() error {
			return httpServer.ListenAndServe()
		},
	)
}
