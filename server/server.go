package server

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
)

type Server interface {
	Start() error
	Shutdown(ctx context.Context) error
}

type server struct {
	echo *echo.Echo
	port string
}

func NewServer(port string) Server {
	return &server{
		port: port,
		echo: echo.New(),
	}
}

func (s *server) Start() error {
	port := fmt.Sprintf(":%s", s.port)

	return s.echo.Start(port)
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}
