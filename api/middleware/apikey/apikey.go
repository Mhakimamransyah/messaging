package apikey

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ApiKey() echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:api-key",
		Validator: func(s string, c echo.Context) (bool, error) {
			return s == os.Getenv("SEROTONIN_KEY"), nil
		},
	})
}
