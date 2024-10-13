package initialize

import (
	"build-service/common/mongodb"
	"build-service/repositories/mongotx"
	"build-service/repositories/user_transaction_history"
)

var (
	repositories *Repositories
)

type Repositories struct {
	IUserTransactionHistoryRepo user_transaction_history.IUserTransactionHistoryRepo
	IMongoTxRepository          mongotx.IMongoTxRepository
}

func NewRepositories(dbStorage *mongodb.DatabaseStorage) *Repositories {
	repositories = &Repositories{
		IUserTransactionHistoryRepo: user_transaction_history.NewRepoUserTransactionHistory(dbStorage),
		IMongoTxRepository:          mongotx.IMongoTxRepository(dbStorage),
	}
	return repositories
}
