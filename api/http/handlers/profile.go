package handlers

import (
	"BuildService/api/http/models"
	"BuildService/common/custom/binding"
	"BuildService/internal/services"
	"BuildService/pkg/helpers/adapters"
	"BuildService/pkg/helpers/resp"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
)

type ProfileHandler struct {
	profileService services.IProfileService
}

func NewProfileHandler(profileService services.IProfileService) *ProfileHandler {
	return &ProfileHandler{
		profileService: profileService,
	}
}

func (h *ProfileHandler) GetUserTransactionHistory(c echo.Context) error {
	ctx := c.Request().Context()
	var req models.GetUserTransactionHistoryReq
	if err := binding.GetBinding().Bind(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, resp.BuildErrorResp(resp.ErrDataInvalid, err.Error(), resp.LangEN))
	}

	dataDomain := adapters.AdapterProfile{}.ConvReq2ServUserTransactionHistoryTx(req)
	data, total, err := h.profileService.GetUserHistoryByProfile(ctx, *dataDomain)
	if err != nil {
		return c.JSON(http.StatusNotFound, resp.BuildErrorResp(err.ErrorCode, err.Description, resp.LangEN))
	}

	rs := resp.BuildSuccessResp(resp.LangEN, data)
	rs.Paging = &resp.Paging{
		Total:  total,
		Offset: req.Offset,
		Limit:  req.Limit,
	}

	return c.JSON(http.StatusOK, rs)
}

func (h *ProfileHandler) CreateUserTransactionHistory(c echo.Context) error {
	ctx := c.Request().Context()
	var req models.UserTransactionHistory
	if err := binding.GetBinding().Bind(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, resp.BuildErrorResp(resp.ErrDataInvalid, err.Error(), resp.LangEN))
	}

	dataDomain := adapters.AdapterProfile{}.ConvModelToDomainUserTransactionHistoryTx(req)
	data, err := h.profileService.CreateUserTransactionHistory(ctx, *dataDomain)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create user transaction history")
		return c.JSON(http.StatusNotFound, resp.BuildErrorResp(err.ErrorCode, err.Description, resp.LangEN))
	}
	return c.JSON(http.StatusOK, resp.BuildSuccessResp(resp.LangEN, data))
}

func (h *ProfileHandler) UpdateUserTransactionHistory(c echo.Context) error {
	ctx := c.Request().Context()
	var req models.UserTransactionHistory
	if err := binding.GetBinding().Bind(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, resp.BuildErrorResp(resp.ErrDataInvalid, err.Error(), resp.LangEN))
	}

	dataDomain := adapters.AdapterProfile{}.ConvModelToDomainUserTransactionHistoryTx(req)
	data, err := h.profileService.UpdateUserTransactionHistoryByProfile(ctx, *dataDomain, dataDomain.ProfileID)
	if err != nil {
		return c.JSON(http.StatusNotFound, resp.BuildErrorResp(err.ErrorCode, err.Description, resp.LangEN))
	}
	return c.JSON(http.StatusOK, resp.BuildSuccessResp(resp.LangEN, data))
}

func (h *ProfileHandler) DeleteUserTransactionHistory(c echo.Context) error {
	var req models.GetUserTransactionHistoryByProfileReq
	ctx := c.Request().Context()
	err := h.profileService.DeleteUserTransactionHistoryByProfile(ctx, req.ProfileID)
	if err != nil {
		return c.JSON(http.StatusNotFound, resp.BuildErrorResp(err.ErrorCode, err.Description, resp.LangEN))
	}
	return c.JSON(http.StatusOK, resp.BuildSuccessResp(resp.LangEN, nil))
}
