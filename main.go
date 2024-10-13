package main

import (
	apiHttp "build-service/api/http"
	"build-service/common/logger"
	"build-service/common/mongodb"
	"build-service/config"
	"build-service/initialize"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
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

	// Initialize clients
	clients := initialize.NewClients()

	// Initialize repositories
	repo := initialize.NewRepositories(dbStorage)

	// Initialize services
	service := initialize.NewServices(conf, clients, repo)

	// Initialize handlers
	handler := initialize.NewHandlers(service)

	//go func() {
	//	msgBroker := msgbroker.NewMsgBroker(conf, handler.OrderHandler, handler.CorePointHandler)
	//	msgBroker.Start(ctx)
	//}()

	srv := apiHttp.NewHttpServe(conf, handler.ProfileHandler, handler.PointHandler)
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
