package initialize

import (
	"BuildService/config"
	"BuildService/internal/services"
)

type Services struct {
	profileService services.IProfileService
	pointService   services.IPointService
}

func NewServices(
	config *config.SystemConfig,
	clients *Clients,
	repo *Repositories,
	// redisClient *redis.Client,
	// chain *chain.Client,
) *Services {
	profileService := services.NewProfileService(
		config,
		repo.IMongoTxRepository,
		repo.IUserTransactionHistoryRepo,
	)

	pointService := services.NewPointService(
		config,
		clients.ReceiverClient,
	)

	service := &Services{
		profileService: profileService,
		pointService:   pointService,
	}

	return service
}
