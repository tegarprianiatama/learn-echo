package controllers

import (
	"learn-echo/models"
	"learn-echo/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateNewsController(c echo.Context) error {
	var newsRequest models.News
	c.Bind(&newsRequest)

	err := repositories.AddNews(&newsRequest)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, models.BaseResponse{
		Message: "Successfully added data.",
		Status:  true,
		Data:    nil,
	})
}

func GetNewsController(c echo.Context) error {
	var news []models.News

	err := repositories.GetNews(&news)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Successfully displaying data",
		Status:  true,
		Data:    news,
	})
}
