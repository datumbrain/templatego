package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func AboutHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "about.html", map[string]interface{}{
		"name": "About",
		"msg":  "All about Datum Brain!",
	})
}
