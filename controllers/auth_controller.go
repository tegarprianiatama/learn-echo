package controllers

import (
	"learn-echo/middlewares"
	"learn-echo/models"
	"learn-echo/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterController(c echo.Context) error {
	var userRegister models.User
	c.Bind(&userRegister)

	result, _ := bcrypt.GenerateFromPassword(
		[]byte(userRegister.Password), bcrypt.DefaultCost,
	)
	userRegister.Password = string(result)
	err := repositories.Register(&userRegister)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	var userResponse models.RegisterResponse
	userResponse.Id = userRegister.Id
	userResponse.Name = userRegister.Name
	userResponse.Email = userRegister.Email

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Successfully added data.",
		Status:  true,
		Data:    userResponse,
	})
}

func LoginController(c echo.Context) error {
	var userLogin models.User
	c.Bind(&userLogin)

	user, err := repositories.GetUserByEmail(userLogin.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, models.BaseResponse{
			Message: "Invalid email or password",
			Status:  false,
			Data:    nil,
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, models.BaseResponse{
			Message: "Invalid email or password",
			Status:  false,
			Data:    nil,
		})
	}

	var userResponse models.LoginResponse
	userResponse.Id = user.Id
	userResponse.Name = user.Name
	userResponse.Email = user.Email
	userResponse.Token = middlewares.GenerateJWTToken(
		userResponse.Id,
		userResponse.Name,
	)

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Login successful",
		Status:  true,
		Data:    userResponse,
	})
}

var tokenBlacklist = make(map[string]bool)

func LogoutController(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")

	tokenBlacklist[token] = true

	return c.JSON(http.StatusOK, "Logout successful")
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if tokenBlacklist[token] {
			return c.JSON(http.StatusUnauthorized, "Token has been revoked")
		}
		return next(c)
	}
}
