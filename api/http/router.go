package http

import (
	"build-service/api/http/middlewares"
	"build-service/api/http/routers"
	"build-service/pkg/helpers/resp"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const prefixPath = "build-service"
const healthPath = prefixPath + "/v1/health"

func (app *httpServ) InitRouters(e *echo.Echo) {
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Use(middlewares.AddExtraDataForRequestContext)
	e.Use(middlewares.Logging)

	e.GET(healthPath, func(c echo.Context) error {
		if healthCheck {
			return c.JSON(http.StatusOK, resp.BuildSuccessResp(resp.LangEN, nil))
		}
		return c.JSON(http.StatusInternalServerError, resp.BuildErrorResp(resp.ErrSystem, "", resp.LangEN))
	})

	//profile router
	controller := routers.NewProfileController(e, app.profileHandler)
	controller.SetupProfileRoutes()

	//point router
	pointController := routers.NewPointController(e, app.pointHandler)
	pointController.SetupPointRoutes()
}
