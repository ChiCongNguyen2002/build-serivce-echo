package services

import (
	"BuildService/common/logger"
	"BuildService/config"
	modelsServ "BuildService/internal/domains"
	"BuildService/pkg/helpers/adapters"
	"BuildService/pkg/helpers/resp"
	"BuildService/repositories/mongo_tx"
	"BuildService/repositories/user_transaction_history"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"time"
)

type ProfileService struct {
	conf        *config.SystemConfig
	mongoRepo   mongo_tx.IMongoTxRepository
	profileRepo user_transaction_history.IUserTransactionHistoryRepo
}

type IProfileService interface {
	GetUserHistoryByProfile(ctx context.Context, req modelsServ.GetUserTransactionHistoryReq) ([]modelsServ.UserTransactionHistory, int64, *resp.CustomError)
	CreateUserTransactionHistory(ctx context.Context, order modelsServ.UserTransactionHistory) (modelsServ.UserTransactionHistory, *resp.CustomError)
	UpdateUserTransactionHistoryByProfile(ctx context.Context, order modelsServ.UserTransactionHistory, profileID string) (modelsServ.UserTransactionHistory, *resp.CustomError)
	DeleteUserTransactionHistoryByProfile(ctx context.Context, profileID string) *resp.CustomError
	CreateOrderTransactionPoint(ctx context.Context, order *modelsServ.OrderSuccessEvent) *resp.CustomError
	CompleteTransactionPoint(ctx context.Context, order *modelsServ.EarnPointOrderEvent) *resp.CustomError
}

func NewProfileService(conf *config.SystemConfig, mongoRepo mongo_tx.IMongoTxRepository, profileRepo user_transaction_history.IUserTransactionHistoryRepo) IProfileService {
	return &ProfileService{
		conf:        conf,
		mongoRepo:   mongoRepo,
		profileRepo: profileRepo,
	}
}

func (p ProfileService) GetUserHistoryByProfile(ctx context.Context, req modelsServ.GetUserTransactionHistoryReq) ([]modelsServ.UserTransactionHistory, int64, *resp.CustomError) {
	var profileID string
	var txTypes []string
	var recentMonth time.Time

	if req.ProfileID != "" {
		profileID = req.ProfileID
	}

	if req.RecentMonth > 0 {
		now := time.Now()
		recentMonth = now.AddDate(0, -req.RecentMonth, 0)
	}

	//upper case txType and status
	req.TxType = strings.ToUpper(req.TxType)
	req.Status = strings.ToUpper(req.Status)

	//get info user history
	userHistoryTxs, totalTxs, err := p.profileRepo.GetUserTransactionHistoryByProfile(ctx, profileID, txTypes, recentMonth, req.Offset, req.Limit, req.Status)
	if err != nil {
		return nil, 0, &resp.CustomError{ErrorCode: resp.ErrNotFound, Description: err.Error()}
	}

	pointTxsServ := adapters.AdapterProfile{}.ConvRepo2DomainServArrayUserTransactionHistoryTx(userHistoryTxs)
	return pointTxsServ, totalTxs, nil
}

func (p ProfileService) CreateUserTransactionHistory(ctx context.Context, order modelsServ.UserTransactionHistory) (modelsServ.UserTransactionHistory, *resp.CustomError) {
	pointTxsServ := adapters.AdapterProfile{}.ConvDomainToRepo(order)
	userHistory, err := p.profileRepo.CreateUserTransactionHistory(ctx, &pointTxsServ)
	if err != nil {
		return order, nil
	}
	pointTxsServToDomain := adapters.AdapterProfile{}.ConvRepoToDomain(userHistory)
	return pointTxsServToDomain, nil
}

func (p ProfileService) UpdateUserTransactionHistoryByProfile(ctx context.Context, order modelsServ.UserTransactionHistory, profileID string) (modelsServ.UserTransactionHistory, *resp.CustomError) {
	pointTxsServ := adapters.AdapterProfile{}.ConvDomainToRepo(order)
	userHistory, err := p.profileRepo.UpdateUserTransactionHistoryByProfile(ctx, &pointTxsServ, profileID)
	if err != nil {
		return order, nil
	}
	pointTxsServToDomain := adapters.AdapterProfile{}.ConvRepoToDomain(userHistory)
	return pointTxsServToDomain, nil
}

func (p ProfileService) DeleteUserTransactionHistoryByProfile(ctx context.Context, profileID string) *resp.CustomError {
	err := p.profileRepo.DeleteUserTransactionHistoryByProfile(ctx, profileID)
	if err != nil {
		return nil
	}
	return nil
}

func (s *ProfileService) CreateOrderTransactionPoint(ctx context.Context, order *modelsServ.OrderSuccessEvent) *resp.CustomError {
	log := logger.GetLogger().AddTraceInfoContextRequest(ctx)
	log.Info().Interface("order", order).Msg("CreateOrderTransactionPoint - Start")

	err := s.mongoRepo.ExecTransaction(ctx, func(sessionCtx mongo.SessionContext) (interface{}, error) {
		// Convert order to repo models
		newOrder := &modelsServ.OrderSuccessEvent{}
		newOrder.BuildCreateOrderTransaction(order)
		orderRepoModel := adapters.AdapterProfile{}.ConvertOrderCreateDomainToRepo(newOrder)

		if err := s.profileRepo.UpsertCreateOrderTransaction(sessionCtx, orderRepoModel); err != nil {
			log.Error().Err(err).Msg("CreateOrderTransactionPoint - Failed to upsert order")
			return nil, err
		}

		log.Info().Msg("CreateOrderTransactionPoint - Order upsert successfully")
		return nil, nil
	})

	if err != nil {
		return &resp.CustomError{ErrorCode: resp.ErrSystem, Description: err.Error()}
	}

	return nil
}

func (s *ProfileService) CompleteTransactionPoint(ctx context.Context, order *modelsServ.EarnPointOrderEvent) *resp.CustomError {
	log := logger.GetLogger().AddTraceInfoContextRequest(ctx)
	log.Info().Interface("order", order).Msg("CompleteOrderEarnPoint - Start")

	err := s.mongoRepo.ExecTransaction(ctx, func(sessionCtx mongo.SessionContext) (interface{}, error) {
		orderRepoModel := adapters.AdapterProfile{}.ConvertCompleteOrderDomainToRepo(order)

		if err := s.profileRepo.UpsertCompleteOrderTransaction(sessionCtx, orderRepoModel); err != nil {
			return nil, err
		}
		return nil, nil
	})

	if err != nil {
		log.Error().Err(err).Msg("CompleteOrderEarnPoint - Transaction failed")
		return &resp.CustomError{ErrorCode: resp.ErrSystem, Description: err.Error()}
	}

	log.Info().Msg("CompleteOrderEarnPoint - Order upsert successfully")
	return nil
}
