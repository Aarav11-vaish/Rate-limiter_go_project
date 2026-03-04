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
var customLimiter = NewTockenBucketLimiter(5, 1) // 5 tokens max, refills at 1 token per second
var concurrencySem = make(chan struct{}, 3) // limit to 3 concurrent requests
	e.GET("/weather/:city", func(c echo.Context) error {
    if !customLimiter.allow() {
		return c.JSON(429, map[string]string{"error": "Rate limit exceeded. Please try again later."})
	}
		city := c.Param("city")
		concurrencySem<- struct{}{} // acquire a filed tht means we are processing a request
		defer func() { <-concurrencySem }() // release the field when done
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
