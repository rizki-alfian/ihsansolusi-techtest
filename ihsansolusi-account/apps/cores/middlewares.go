package cores

import (
	"ihsansolusi-account/apps/cores/middlewares"
	"github.com/labstack/echo/v4"
)

func RegisterMiddlewares(e *echo.Echo) {
	middlewares.InitLogger()
	e.Use(middlewares.DefaultLogger())
}