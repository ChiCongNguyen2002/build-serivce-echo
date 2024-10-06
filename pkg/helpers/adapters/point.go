package adapters

import (
	modelsHandler "BuildService/api/http/models"
	modelsServ "BuildService/internal/domains"
)

type AdapterLPPoint struct{}

func (a AdapterLPPoint) ConvertOrderHandler2Domain(d modelsHandler.OrderRequest) (data modelsServ.Order) {
	return modelsServ.Order{
		OrderNumber: d.OrderNumber,
		CreateTime:  d.CreateTime,
		Amount:      d.Amount,
		Currency:    d.Currency,
		VGAUserID:   d.VGAUserID,
		SourceType:  d.SourceType,
	}
}
