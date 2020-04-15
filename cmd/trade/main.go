package main

import (
	"os"
	"os/signal"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/jesseobrien/trade/internal/httpsrv"
	"github.com/jesseobrien/trade/internal/service"
)

func main() {

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	defer signal.Stop(quit)

	logger := log.Logger{
		Handler: cli.New(os.Stdout),
	}

	logger.Info("📈 Welcome to Trade 📈")

	market := service.NewMarket(logger)

	go market.Run()

	httpSrv := httpsrv.NewHTTPServer(logger)

	go httpSrv.Run()

	<-quit

}
