package middleware

import (
	"github.com/ken-aio/go-echo-sqlboiler/app/infra/db"
	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
)

// Init initialize database connection
func Init(e *echo.Echo) {
	db.Init()

	e.Use(echoMw.Logger())
	e.Use(echoMw.Recover())

	e.Use(ClientTrace)
}
