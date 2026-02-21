package main

import (
	"fmt"

	fetch_weather_data "github.com/Aarav11-vaish/Rate-limiter_go_project/fetch_weather_data"
	"github.com/labstack/echo/v4"
	"github.com/joho/godotenv"
)

func main() {
	e := echo.New()
    if err := godotenv.Load("../.env"); err != nil {
        fmt.Println("Warning: Could not load .env file:", err)
    }

	// print something
	fmt.Println("Welcome to the server backend oF GO hit /weather:{city name} to get the details and / to get a welcome message")

	e.GET("/weather/:city", func(c echo.Context) error {
		city := c.Param("city")
		weatherData, err := fetch_weather_data.FetchWeatherData(city)
		if err != nil {
			return c.JSON(500, map[string]string{"error": "Failed to fetch weather data"})
		}
		return c.JSON(200, weatherData)
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Welcome to the server backend oF GO hit /getIP to get your IP address with region and country and / to get a welcome message")
	})
	e.Start(":8080")
}
