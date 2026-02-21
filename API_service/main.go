package main

import (
	"fmt"
     "github.com/Aarav11-vaish/rate-limiter_go_roject/rate-limiter/API_service/geo_location_through_IP"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// print something
	fmt.Println("Welcome to the server backend oF GO hit /getIP to get your IP address with region and country and / to get a welcome message")

	// wnat to use mvc pattern
	e.GET("/getIP", func(c echo.Context) error {

		// get the ip address of the client

		ip := c.RealIP()
		// get the region and country of the ip address

		data, err = getlocation(ip)
		if err != nil {
			return c.String(500, "Error getting location")
		}
		return c.String(200, data)

	})

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Welcome to the server backend oF GO hit /getIP to get your IP address with region and country and / to get a welcome message")
	})
	e.Start(":8080")
}
