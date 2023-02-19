package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func Signup(c echo.Context) error {
	c.Logger().Debug("signup-----------------------------------------")
	c.Logger().Debug(c.ParamNames())
	return c.String(http.StatusCreated, "Signup! Content")
}

func Login(c echo.Context) error {
	c.Logger().Debug("login----------------------------------------")
	return c.String(http.StatusOK, "Login! Content")
}
