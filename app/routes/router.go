package routes

import (
	"net/http"
	"time"

	"github.com/ken-aio/go-echo-sqlboiler/app/handlers/v1"
	"github.com/ken-aio/go-echo-sqlboiler/app/middleware"
	"github.com/labstack/echo"

	"github.com/swaggo/echo-swagger"
)

// Init initialize echo application
func Init() *echo.Echo {
	e := echo.New()

	e.Debug = true

	middleware.Init(e)

	// swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/routes", func(c echo.Context) error {
		return c.JSON(http.StatusOK, e.Routes())
	})
	// ping
	e.GET("/ping", func(c echo.Context) error {
		resp := map[string]time.Time{"pong": time.Now()}
		return c.JSON(http.StatusOK, resp)
	})

	apiV1 := e.Group("api/v1")
	{
		users := apiV1.Group("/users")
		{
			users.POST("", v1.UserCreate)
			users.GET("", v1.UserList)
			users.PUT("/:id", v1.UserUpdate)
			users.DELETE("/:id", v1.UserDelete)
		}
		groups := apiV1.Group("/groups")
		{
			mems := groups.Group("/:group_id/members")
			{
				mems.GET("", v1.GroupMemberList)
			}
		}
	}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	return e
}
