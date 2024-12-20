package router

import (
	"github.com/labstack/echo/v4"
	"github.com/yowger/pet-day-care-api-2/config"
)

func Init(e *echo.Echo) {
	setCORSConfig(e)

	setUserController(e)
	setPetController(e)
	setBookingController(e)
}

func setCORSConfig(e *echo.Echo) {
}

func setUserController(e *echo.Echo) {
	apiGroup := e.Group(config.API)
	userGroup := apiGroup.Group(config.APIUserGroup)

	// userGroup.GET("/", )
}

func setPetController(e *echo.Echo) {
}

func setBookingController(e *echo.Echo) {
}
