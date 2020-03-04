package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/andodevel/go-echo-template/helpers"
)

var jwtMiddlewareFnc echo.MiddlewareFunc

func init() {
	config := middleware.DefaultJWTConfig
	config.TokenLookup = "cookie:" + helpers.JWTCookieKey
	config.SigningKey = []byte(helpers.JWTSecrect)
	jwtMiddlewareFnc = middleware.JWTWithConfig(config)
}

// JWT ...
func JWT() echo.MiddlewareFunc {
	return jwtMiddlewareFnc
}
