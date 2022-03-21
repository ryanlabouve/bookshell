package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	bookshell "ryanlabouve.com/bookshell/lib"
)

type Noop struct{}

func main() {
	bookshell.Load()
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/hotdog", func(c echo.Context) (err error) {
		error := bookshell.Load()

		if error != nil {
			echo.NewHTTPError(http.StatusTeapot, "Boop")
		}

		return c.JSON(http.StatusOK, Noop{})
	})

	e.Logger.Fatal(e.Start(":1323"))
}
