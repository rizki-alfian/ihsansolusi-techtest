package middlewares

import (
	"time"
	"os"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func InitLogger() {
	env := os.Getenv("APP_ENV")
	
	logger, _ := zap.NewProduction()
	if env != "production" {
		// logger, _ = zap.NewDevelopment()
	}

	defer logger.Sync()
	Logger = logger.Sugar()
}

func RequestLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		if err != nil {
			c.Error(err)
		}

		status := c.Response().Status
		duration := time.Since(start)

		if status >= 500 {
			Logger.Errorw("SERVER ERROR",
				"method", c.Request().Method,
				"path", c.Request().URL.Path,
				"status", status,
				"duration", duration.String(),
			)
		} else if status >= 400 {
			Logger.Warnw("CLIENT ERROR",
				"method", c.Request().Method,
				"path", c.Request().URL.Path,
				"status", status,
				"duration", duration.String(),
			)
		} else {
			Logger.Infow("Request processed",
				"method", c.Request().Method,
				"path", c.Request().URL.Path,
				"status", status,
				"duration", duration.String(),
			)
		}

		return err
	}
}

func DefaultLogger() echo.MiddlewareFunc {
	return RequestLogger
}