package initialize

import (
	"BuildService/common/mongodb"
	"BuildService/repositories/mongo_tx"
	"BuildService/repositories/user_transaction_history"
)

var (
	repositories *Repositories
)

type Repositories struct {
	IUserTransactionHistoryRepo user_transaction_history.IUserTransactionHistoryRepo
	IMongoTxRepository          mongo_tx.IMongoTxRepository
}

func NewRepositories(dbStorage *mongodb.DatabaseStorage) *Repositories {
	repositories = &Repositories{
		IUserTransactionHistoryRepo: user_transaction_history.NewRepoUserTransactionHistory(dbStorage),
		IMongoTxRepository:          mongo_tx.IMongoTxRepository(dbStorage),
	}
	return repositories
}
