package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func HealtCheckHandler(c echo.Context)error{
	return c.String(http.StatusOK, "OK!")
}
