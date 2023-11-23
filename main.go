package main

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pharrisee/htmxample/views"
)

// using Echo but any framework would be very similar

func main() {
	e := echo.New()
	// add middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return views.Index().Render(context.Background(), c.Response().Writer)
	})

	log.Fatalln(e.Start(":2112"))
}
