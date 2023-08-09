package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"middleWare/internal/app/endpoint"
	"middleWare/internal/app/middleWare"
	"middleWare/internal/app/service"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.Use(middleWare.AdminCheck)

	a.echo.GET("/test", a.e.Status)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("Server is working")

	err := a.echo.Start(":8080")

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
