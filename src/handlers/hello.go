package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Hello(ctx echo.Context) (err error) {
	response := "Hello World!"
	return ctx.JSON(http.StatusOK, response)
}
