package handlers

import (
	"build-service/api/http/models"
	"build-service/common/custom/binding"
	"build-service/internal/services"
	"build-service/pkg/helpers/adapters"
	"build-service/pkg/helpers/resp"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PointHandler struct {
	pointService services.IPointService
}

func NewPointHandler(pointService services.IPointService) *PointHandler {
	return &PointHandler{
		pointService: pointService,
	}
}

func (h *PointHandler) CreatePointTransaction(c echo.Context) error {
	ctx := c.Request().Context()
	var req *models.OrderRequest
	if err := binding.GetBinding().Bind(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, resp.BuildErrorResp(resp.ErrDataInvalid, err.Error(), resp.LangEN))
	}
	dataDomain := adapters.AdapterLPPoint{}.ConvertOrderHandler2Domain(req)
	err := h.pointService.CreatePointTransaction(ctx, dataDomain)
	if err != nil {
		return c.JSON(http.StatusNotFound, resp.BuildErrorResp(err.ErrorCode, err.Description, resp.LangEN))
	}
	return c.JSON(http.StatusOK, resp.BuildSuccessResp(resp.LangEN, nil))
}
