package routers

import (
	"BuildService/api/http/handlers"
	"github.com/labstack/echo/v4"
)

type PointController struct {
	e         *echo.Echo
	clientSys *echo.Group
	handlers  *handlers.PointHandler
}

func NewPointController(e *echo.Echo, handlers *handlers.PointHandler) *PointController {
	return &PointController{
		e:         e,
		clientSys: e.Group(prefixSystemPath),
		handlers:  handlers,
	}
}

func (app *PointController) SetupPointRoutes() {
	// Set up router for point
	app.SetupRouterPoint()
}

func (app *PointController) SetupRouterPoint() {
	profile := app.clientSys.Group(prefixPoint)
	profile.POST(prefixPointTransactionPath, app.handlers.CreatePointTransaction)
}
