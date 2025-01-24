package controllers

import (
	"mentor_hub-api/responces"
	"net/http"

	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	return c.JSON(http.StatusOK, responces.Login{
		Status:  http.StatusOK,
		Message: "Login Sucussful done",
		Data:    nil,
	})
}
