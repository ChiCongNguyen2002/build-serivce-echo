package http

import (
	"BuildService/api/http/routers"
	"BuildService/pkg/helpers/resp"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

const prefixPath = "build-service"
const healthPath = prefixPath + "/v1/health"

func (app *httpServ) InitRouters(e *echo.Echo) {
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	//e.Use(middlewares.AddExtraDataForRequestContext)
	//e.Use(middlewares.Logging)
	//e.Use(middlewares.Region)

	e.GET(healthPath, func(c echo.Context) error {
		if healthCheck {
			return c.JSON(http.StatusOK, resp.BuildSuccessResp(resp.LangEN, nil))
		}
		return c.JSON(http.StatusInternalServerError, resp.BuildErrorResp(resp.ErrSystem, "", resp.LangEN))
	})

	//point router
	controller := routers.NewProfileController(e, app.pointHandler)
	controller.SetupProfileRoutes()
}
