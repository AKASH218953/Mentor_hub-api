package main

import (
	"mentor_hub-api/rotues"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS default
	//Allows requests from any origin wth GET, HEAD, PUT, POST or DELETE method.
	e.Use(middleware.CORS())
	//CORS restricted
	//Allows requests from any `https://labstack.com` or `https://labstack.net` origin
	//wth GET, PUT, POST or DELETE method.
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"https://localhost:45615"}, // allow origin from frontend
	// 	AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	// }))
	//run database

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, &echo.Map{"data": "Hello from MentorHub"})
	})

	rotues.Login(e)

	e.Logger.Fatal(e.Start(":8080"))
}
