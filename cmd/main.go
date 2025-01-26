package main

import (
	"log"
	"net/http"
	"template/view"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	app.Renderer = view.New()

	app.Use(middleware.Logger())
	app.Static("/static", "public/static")

	app.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", echo.Map{})
	})

	app.POST("/", func(c echo.Context) error {
		log.Println(c.FormValue("counter"))
		return c.Redirect(303, "/")
	})

	if err := app.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
