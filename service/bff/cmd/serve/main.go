package main

import (
	"context"
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

	restServer.RegisterRoute()

	go func() {
		if err := restServer.Serve(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := restServer.Shutdown(ctx); err != nil {
		panic(err)
	}
}
