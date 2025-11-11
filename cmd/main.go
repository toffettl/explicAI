package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/toffettl/explicAI/configuration"
	"github.com/toffettl/explicAI/internal/infrastructure/log"
)

func main() {
	c := configuration.Init()
	go configuration.NewApplication(c).Start()
	shutDown()
}

func shutDown() {
	signalShutDown := make(chan os.Signal, 2)
	signal.Notify(signalShutDown, syscall.SIGINT, syscall.SIGTERM)
	switch <-signalShutDown {
	case syscall.SIGINT:
		log.LogInfo(context.Background(), "SIGINT signal, explicaAI is stopping...")
	case syscall.SIGTERM:
		log.LogInfo(context.Background(), "SIGTERM signal, explicaAI is stopping...")
	}
}
