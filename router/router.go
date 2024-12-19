package router

import "github.com/labstack/echo/v4"

func Init(e *echo.Echo) {
	setCORSConfig(e)

	setUserController(e)
	setPetController(e)
	setBookingController(e)
}

func setCORSConfig(e *echo.Echo) {
}

func setUserController(e *echo.Echo) {
}

func setPetController(e *echo.Echo) {
}

func setBookingController(e *echo.Echo) {
}
