package initialize

import (
	"BuildService/api/http/handlers"
	"BuildService/api/msgbroker/consumer/core_handle_point"
	"BuildService/api/msgbroker/consumer/order"
)

type Handlers struct {
	ProfileHandler   *handlers.ProfileHandler
	OrderHandler     *order.OrderHandler
	CorePointHandler *core_handle_point.CorePointHandler
}

func NewHandlers(services *Services) *Handlers {

	profileHandler := handlers.NewProfileHandler(
		services.profileService,
	)

	orderHandler := order.NewOrderHandler(
		services.profileService,
	)

	corePointHandler := core_handle_point.NewCorePointHandler(
		services.profileService,
	)

	return &Handlers{
		ProfileHandler:   profileHandler,
		OrderHandler:     orderHandler,
		CorePointHandler: corePointHandler,
	}
}
