package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"name": "Home",
		"msg":  "Hello, Datum Brain!",
	})
}
