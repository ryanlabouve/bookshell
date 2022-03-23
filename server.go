package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	bookshell "ryanlabouve.com/bookshell/lib"
)

type Noop struct{}

func main() {
	bookshell.Load()
	db := bookshell.InitializeDb()
	bookshell.SeedDb(db)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/api/books", func(c echo.Context) error {
		return c.JSON(http.StatusOK, bookshell.GetBooks(db))
	})

	e.POST("/hotdog", func(c echo.Context) (err error) {
		error := bookshell.Load()

		if error != nil {
			echo.NewHTTPError(http.StatusTeapot, "Boop")
		}

		return c.JSON(http.StatusOK, Noop{})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
