package server

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
)

type EchoServer struct {
	echo *echo.Echo
	port string
}

func NewServer(port string) *EchoServer {
	return &EchoServer{
		port: port,
		echo: echo.New(),
	}
}

func (s *EchoServer) Start() error {
	port := fmt.Sprintf(":%s", s.port)

	return s.echo.Start(port)
}

func (s *EchoServer) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}
