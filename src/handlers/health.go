package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Health(ctx echo.Context) (err error) {
	response := "OK"
	return ctx.JSON(http.StatusOK, response)
}
