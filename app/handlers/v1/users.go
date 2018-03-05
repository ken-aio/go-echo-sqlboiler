package v1

import (
	"net/http"

	"github.com/ken-aio/go-echo-sqlboiler/app/models"
	"github.com/labstack/echo"
)

// UserCreate User create API
// @Summary User create API
// @Description create new user
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {object} models.Test	"create ok"
// @Router /api/v1/users [get]
func UserCreate(c echo.Context) error {
	res := models.Test{"hoge"}
	return c.JSON(http.StatusOK, res)
}

func UserList(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"list": "ok"})
}

func UserUpdate(c echo.Context) error {
	return c.JSON(http.StatusOK, "update ok")
}

func UserDelete(c echo.Context) error {
	return c.JSON(http.StatusOK, "delete ok")
}
