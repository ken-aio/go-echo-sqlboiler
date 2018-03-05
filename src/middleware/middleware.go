package middleware

import (
	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
)

func Init(e *echo.Echo) {
	e.Use(echoMw.Logger())
	e.Use(echoMw.Recover())
}
