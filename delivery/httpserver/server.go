package httpserver

import (
	"clean-code-structure/config"
	"clean-code-structure/delivery/httpserver/healthhandler"
	"clean-code-structure/logger"
	"clean-code-structure/service/healthservice"
	"clean-code-structure/validator/healthvalidator"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type Server struct {
	config             config.Config
	Router             *echo.Echo
	healthcheckHandler healthhandler.Handler
}

type RequiredServices struct {
	HealthService healthservice.Service
}

type RequiredValidators struct {
	HealthValidator healthvalidator.Validator
}

func New(config config.Config, services RequiredServices, validators RequiredValidators) Server {
	return Server{
		Router:             echo.New(),
		config:             config,
		healthcheckHandler: healthhandler.New(services.HealthService, validators.HealthValidator),
	}
}

func (s Server) Serve() {
	// Middleware

	s.Router.Use(middleware.RequestID())

	s.Router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:           true,
		LogStatus:        true,
		LogHost:          true,
		LogRemoteIP:      true,
		LogRequestID:     true,
		LogMethod:        true,
		LogContentLength: true,
		LogResponseSize:  true,
		LogLatency:       true,
		LogError:         true,
		LogProtocol:      true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			errMsg := ""
			if v.Error != nil {
				errMsg = v.Error.Error()
			}

			logger.Logger.Named("http-server").Info("request",
				zap.String("request_id", v.RequestID),
				zap.String("host", v.Host),
				zap.String("content-length", v.ContentLength),
				zap.String("protocol", v.Protocol),
				zap.String("method", v.Method),
				zap.Duration("latency", v.Latency),
				zap.String("error", errMsg),
				zap.String("remote_ip", v.RemoteIP),
				zap.Int64("response_size", v.ResponseSize),
				zap.String("uri", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))

	s.Router.Use(middleware.Recover())

	// Routes
	s.healthcheckHandler.SetRoutes(s.Router.Group("health"))

	// Start server
	address := fmt.Sprintf(":%d", s.config.HTTPServer.Port)
	fmt.Printf("start echo server on %s\n", address)
	if err := s.Router.Start(address); err != nil {
		fmt.Println("router start error", err)
	}
}
