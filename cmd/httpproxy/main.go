package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/qdm12/gluetun/internal/httpproxy"
	"github.com/qdm12/golibs/logging"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	const stealth = false
	const verbose = true
	const username = "user"
	const password = "password"
	const address = ":8000"

	logger := logging.New(logging.Settings{})

	server := httpproxy.New(ctx, address, logger, stealth, verbose, username, password)

	errorCh := make(chan error)
	go server.Run(ctx, errorCh)

	select {
	case <-ctx.Done():
		logger.Info("done")
		stop()
	case err := <-errorCh:
		logger.Error(err.Error())
	}
}
