package main

import (
	"context"
	"github.com/bosskrub9992/templatebe/service/bff/src/router"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	restServer, closeRestServer, err := InitializeRestServer()
	if err != nil {
		panic(err)
	}
	defer closeRestServer()

	router.RegisterRoute(restServer)

	go func() {
		if err := restServer.Serve(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, os.Interrupt)

	<-gracefulShutdown

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := restServer.Shutdown(ctx); err != nil {
		panic(err)
	}
}
