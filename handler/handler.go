package handler

import (
	"log"

	"github.com/labstack/echo/v4"
)

func Signup(c echo.Context) error {
	log.Println("signup")
	return nil
}

func Login(c echo.Context) error {
	log.Println("login")
	return nil
}
