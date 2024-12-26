package router

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/yowger/pet-day-care-api-2/config"
	"github.com/yowger/pet-day-care-api-2/controller"
	sqlc "github.com/yowger/pet-day-care-api-2/database/sqlc"
	"github.com/yowger/pet-day-care-api-2/repository"
	"github.com/yowger/pet-day-care-api-2/service"
)

func NewRouter(e *echo.Echo, queries *sqlc.Queries, ctx context.Context) {

	setUpUserRoutes(e, queries, ctx)
}

func setUpUserRoutes(e *echo.Echo, queries *sqlc.Queries, ctx context.Context) {
	userRepo := repository.NewUserRepo(queries)
	userService := service.NewUserService(userRepo, ctx)
	userController := controller.NewUserController(userService)

	e.GET(config.APIUser, func(c echo.Context) error { return userController.GetUser(c) })
}
