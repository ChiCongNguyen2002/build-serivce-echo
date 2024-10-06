package http

import (
	"BuildService/api/http/handlers"
	"BuildService/common/logger"
	"BuildService/config"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

var (
	healthCheck bool
	mu          sync.RWMutex
)

func SetHealthCheck(status bool) {
	mu.Lock()
	defer mu.Unlock()
	healthCheck = status
}

type HttpServInterface interface {
	Start(e *echo.Echo)
}

type httpServ struct {
	conf           *config.SystemConfig
	profileHandler *handlers.ProfileHandler
	pointHandler   *handlers.PointHandler
	//coreHandler    *order.OrderHandler
	//earnHandler    *core_handle_point.CorePointHandler
}

func NewHttpServe(
	conf *config.SystemConfig,
	profileHandler *handlers.ProfileHandler,
	pointHandler *handlers.PointHandler,
	// coreHandler *order.OrderHandler,
	// earnHandler *core_handle_point.CorePointHandler,
) *httpServ {
	return &httpServ{
		conf:           conf,
		profileHandler: profileHandler,
		pointHandler:   pointHandler,
		//coreHandler:    coreHandler,
		//earnHandler:    earnHandler,
	}
}

func (app *httpServ) Start(e *echo.Echo) {
	log := logger.GetLogger()
	app.InitRouters(e)
	httpPort := app.conf.HttpPort
	go func() {
		err := e.Start(fmt.Sprintf(":%d", httpPort))
		if err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("can't start echo")
		}
	}()
	log.Info().Msg("all services already")
}
