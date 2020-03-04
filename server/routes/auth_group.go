package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/andodevel/clock_server/server/handlers"
	"github.com/andodevel/clock_server/server/middlewares"
)

// AuthGroup ...
func AuthGroup(e *echo.Echo) {
	g := e.Group("/auth")
	{
		g.POST("/register", handlers.Register())
		g.GET("/login", handlers.LoginView())
		g.POST("/login", handlers.Login())
		g.GET("/me", handlers.CurrentUser())
		g.GET("/logout", handlers.Logout())
		g.POST("/logout", handlers.Logout())
	}
	g.Group("/me", middlewares.JWT()).GET("", handlers.CurrentUser())
}
