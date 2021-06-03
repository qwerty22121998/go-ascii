package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/qwerty22121998/go-ascii/controller"
	"github.com/qwerty22121998/go-ascii/service"
	"net/http"
	"os"
)

type Server struct {
	e          *echo.Echo
	service    *service.Provider
	controller *controller.Provider
}

func NewServer() *Server {
	serviceProvider := service.NewProvider()
	controllerProvider := controller.NewProvider(serviceProvider)

	return &Server{
		e:          echo.New(),
		service:    serviceProvider,
		controller: controllerProvider,
	}
}

func (s *Server) Start() {
	s.e.Use(middleware.Logger())
	s.e.Validator = NewValidator()
	s.e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	s.e.Logger.Fatal(s.e.Start(":" + os.Getenv("PORT")))
}
