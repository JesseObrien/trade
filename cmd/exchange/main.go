package main

import (
	"os"
	"os/signal"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/jesseobrien/trade/internal/exchange"
	"github.com/jesseobrien/trade/internal/nats"
)

func main() {

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	defer signal.Stop(quit)

	logger := log.Logger{
		Handler: cli.New(os.Stdout),
	}

	natsConn, err := nats.NewJSONConnection()

	if err != nil {
		logger.Errorf("could not connect to nats %v", err)
		panic(err)
	}

	logger.Info("📈 Welcome to Trade 📈")

	ex := exchange.New(logger, natsConn)

	go ex.Run()

	<-quit
}
