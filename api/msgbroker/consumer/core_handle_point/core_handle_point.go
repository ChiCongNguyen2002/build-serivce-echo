package core_handle_point

import (
	"BuildService/api/msgbroker/models"
	"BuildService/common/logger"
	"BuildService/internal/services"
	"BuildService/pkg/helpers/adapters"
	"context"
	"errors"
)

type CorePointHandler struct {
	profileService services.IProfileService
}

func NewCorePointHandler(profileService services.IProfileService) *CorePointHandler {
	return &CorePointHandler{
		profileService: profileService,
	}
}

func (h *CorePointHandler) CorePointHandle(ctx context.Context, data models.EarnPointOrderEvent) error {
	log := logger.GetLogger().AddTraceInfoContextRequest(ctx)
	dataServ := adapters.AdapterProfile{}.ConvertEventEarnPointToDomain(&data)
	if err := h.profileService.CompleteTransactionPoint(ctx, dataServ); err != nil {
		log.Err(err).Msg("Failed to complete order and earn points")
		return errors.New(err.Error())
	}
	log.Info().Msg("Earn point order event processed successfully")
	return nil
}
