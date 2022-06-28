package fkBootstrap

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Shutdowns map[string]func() error

func GracefulShutdown(ctx context.Context, funcs Shutdowns) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGABRT, syscall.SIGQUIT)

	sig := <-sigs
	e, _ := errgroup.WithContext(ctx)
	for _, value := range funcs {
		e.Go(value)
	}
	if err := e.Wait(); err != nil {
		log.Panic(err)
	}
	fmt.Println("Exiting server on ", sig)
	os.Exit(0)
}
