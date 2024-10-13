package routers

import (
	"build-service/api/http/handlers"

	"github.com/labstack/echo/v4"
)

type ProfileController struct {
	e         *echo.Echo
	clientSys *echo.Group
	handlers  *handlers.ProfileHandler
}

func NewProfileController(e *echo.Echo, handlers *handlers.ProfileHandler) *ProfileController {
	return &ProfileController{
		e:         e,
		clientSys: e.Group(prefixSystemPath),
		handlers:  handlers,
	}
}

func (app *ProfileController) SetupProfileRoutes() {
	// Set up router for profile
	app.SetupRouterProfile()
}

func (app *ProfileController) SetupRouterProfile() {
	profile := app.clientSys.Group(prefixProfile)
	profile.GET(prefixUserTransactionHistoryPath, app.handlers.GetUserTransactionHistory)
	profile.POST(prefixUserTransactionHistoryPath, app.handlers.CreateUserTransactionHistory)
	profile.PUT(prefixUserTransactionHistoryPath, app.handlers.UpdateUserTransactionHistory)
	profile.DELETE(prefixUserTransactionHistoryPath, app.handlers.DeleteUserTransactionHistory)
}
