package initialize

import (
	"BuildService/config"
	"BuildService/internal/services"
)

type Services struct {
	profileService services.IProfileService
}

func NewServices(
	config *config.SystemConfig,
	repo *Repositories,
	// redisClient *redis.Client,
	// chain *chain.Client,
) *Services {
	profileService := services.NewProfileService(
		config,
		repo.IMongoTxRepository,
		repo.IUserTransactionHistoryRepo,
	)
	service := &Services{
		profileService: profileService,
	}

	return service
}
