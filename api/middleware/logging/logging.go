package logging

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Logging() echo.MiddlewareFunc {
	// Middleware
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status} error=${error} path=${path} time=${time_unix} ip=${ip} \n",
		Output: os.Stdout,
	})
}
