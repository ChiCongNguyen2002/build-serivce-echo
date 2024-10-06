package initialize

import (
	"BuildService/common/logger"
	"BuildService/common/mongodb"
	"BuildService/config"
	"context"
)

var (
	databaseConnection *DatabaseConnection
)

type DatabaseConnection struct {
	Conn *mongodb.DatabaseStorage
}

func NewDatabaseConnection(ctx context.Context) *DatabaseConnection {
	log := logger.GetLogger().AddTraceInfoContextRequest(ctx)

	conn, err := mongodb.ConnectMongoDB(ctx, &config.GetInstance().MongoDBConfig)
	handleError(log, err, "Connect MongoDB promotion failed!")

	databaseConnection = &DatabaseConnection{
		Conn: conn,
	}

	return databaseConnection
}

func handleError(log *logger.Logger, err error, errMsg string) {
	if err != nil {
		log.Fatal().Msgf(errMsg, err)
	}
}