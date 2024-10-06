package main

import (
	apiHttp "BuildService/api/http"
	msgbroker "BuildService/api/msgbroker/consumer"
	"BuildService/common/logger"
	"BuildService/common/mongodb"
	"BuildService/config"
	"BuildService/initialize"
	"context"
	"github.com/labstack/echo/v4"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger.InitLog(os.Getenv("SERVICE_ID"))
	log := logger.GetLogger()
	log.Info().Any("service", os.Getenv("SERVICE_ID")).Msg("Start services")

	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Msgf("load config fail! %s", err)
	}

	apiHttp.SetHealthCheck(true)
	e := echo.New()

	ctx, _ := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt, syscall.SIGTERM)

	dbStorage, err := mongodb.ConnectMongoDB(context.Background(), &conf.MongoDBConfig)
	if err != nil {
		log.Fatal().Msgf("connect mongodb failed! %s", err)
	}

	// Initialize repositories
	repo := initialize.NewRepositories(dbStorage)

	// Initialize services
	service := initialize.NewServices(conf, repo)

	// Initialize handlers
	handler := initialize.NewHandlers(service)

	go func() {
		msgBroker := msgbroker.NewMsgBroker(conf, handler.OrderHandler, handler.CorePointHandler)
		msgBroker.Start(ctx)
	}()

	srv := apiHttp.NewHttpServe(conf, handler.ProfileHandler)
	srv.Start(e)

	// handle graceful shutdown
	<-ctx.Done()
	apiHttp.SetHealthCheck(false)
	cancelCtx, cc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cc()
	if err = e.Shutdown(cancelCtx); err != nil {
		log.Fatal().Msgf("force shutdown services")
	}
}
