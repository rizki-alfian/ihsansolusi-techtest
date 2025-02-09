package routes

import (
	"ihsansolusi-account/apps/users"
	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(e *echo.Echo, userHandler *users.UserHandler) {
	e.POST("/user/register", userHandler.RegisterUser)
	e.GET("/user/balance/:account_number", userHandler.CheckBalance)
}