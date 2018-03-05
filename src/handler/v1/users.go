package v1

import (
	"net/http"

	"github.com/labstack/echo"
)

func UserCreate(c echo.Context) error {
	return c.JSON(http.StatusOK, "create ok")
}

func UserList(c echo.Context) error {
	return c.JSON(http.StatusOK, "list ok")
}

func UserUpdate(c echo.Context) error {
	return c.JSON(http.StatusOK, "update ok")
}

func UserDelete(c echo.Context) error {
	return c.JSON(http.StatusOK, "delete ok")
}
