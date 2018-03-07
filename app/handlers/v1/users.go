package v1

import (
	"net/http"
	"time"

	"github.com/ken-aio/go-echo-sqlboiler/app/services"
	"github.com/labstack/echo"
)

type (
	userCreateReq struct {
		Name      string    `json:"name"`
		Birthdate time.Time `json:"birthdate"`
		Gender    string    `json:"gender"`
	}
)

// UserCreate User create API
// @Summary User create API
// @Description create new user
// @Accept  json
// @Produce  json
// @Param   body     body    v1.userCreateReq     true        "user create parameter"
// @Success 200 {object} models.UserCreate	""
// @Router /api/v1/users [post]
func UserCreate(c echo.Context) error {
	u := &userCreateReq{}
	if err := c.Bind(u); err != nil {
		return err
	}
	userID, err := services.UserCreate(u.Name, u.Birthdate, u.Gender)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]uint64{"user_id": userID})
}

// UserList User list API
// @Summary User list API
// @Description list users
// @Accept  json
// @Produce  json
// @Success 200 {object} models.UserList	""
// @Router /api/v1/users [get]
func UserList(c echo.Context) error {
	res, err := services.UserList()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

// UserUpdate User update API
// @Summary User update API
// @Description update users
// @Accept  json
// @Produce  json
// @Param   user_id     path    int     true        "user id parameter"
// @Success 200 string string	""
// @Router /api/v1/users/{user_id} [put]
func UserUpdate(c echo.Context) error {
	return c.JSON(http.StatusOK, "update ok")
}

// UserDelete User delete API
// @Summary User delete API
// @Description delete users
// @Accept  json
// @Produce  json
// @Param   user_id     path    int     true        "user id parameter"
// @Success 200 string string	""
// @Router /api/v1/users/{user_id} [delete]
func UserDelete(c echo.Context) error {
	return c.JSON(http.StatusOK, "delete ok")
}
