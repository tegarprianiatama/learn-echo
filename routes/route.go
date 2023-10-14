package routes

import (
	"learn-echo/controllers"
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo) {
	e.POST("/register", controllers.RegisterController)
	e.POST("/login", controllers.LoginController)

	e.Use(middleware.Logger())
	eAuth := e.Group("")
	eAuth.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))
	eAuth.Use(controllers.AuthMiddleware)

	eAuth.POST("/create-news", controllers.CreateNewsController)
	eAuth.GET("/news", controllers.GetNewsController)
	eAuth.POST("/logout", controllers.LogoutController)
}
