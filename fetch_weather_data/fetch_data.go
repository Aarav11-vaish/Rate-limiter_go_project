package geolocationthroughipgo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Weather struct {
	Description string `json:"description"`
}

type openWeatherAPIResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []Weather `json:"weather"`
}

type WeatherResponse struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
	Description string  `json:"description"`
}

func fetchFromAPI(city string, apiKey string) ([]byte, error) {
	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric",
		city,
		apiKey,
	)
	if url == "" {
		return nil, fmt.Errorf("failed to construct API URL")
	}
	fmt.Println("Constructed URL:", url)
	// what will url consist of ? ans: The url variable will consist of a formatted string that includes the city name and the API key. It will look something like this: "https://api.openweathermap.org/data/2.5/weather?q=CityName&appid=YourAPIKey&units=metric". The actual values for CityName and YourAPIKey will be replaced with the values passed to the function when it is called.
	data, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from API: %v", err)
	}
	defer data.Body.Close()
	body, err := io.ReadAll(data.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read API response: %v", err)
	}
	defer data.Body.Close()

	return body, nil
}

func FetchWeatherData(city string) (*WeatherResponse, error) {
	apiKey := os.Getenv("API_KEY") // Load the API key from environment variable
if apiKey == "" {
        return nil, fmt.Errorf("API_KEY environment variable not set")
    }
	//what does sprintf do ? ans: fmt.Sprintf is a function in the fmt package that formats a string according to a format specifier and returns the resulting string. It works similarly to printf, but instead of printing the formatted string to standard output, it returns it as a string value. In this case, we use fmt.Sprintf to construct the URL for the API request by inserting the city name and API key into the URL template.

	data, err := fetchFromAPI(city, apiKey)

	if err != nil {
		return nil, err
	}

	var apiResponse openWeatherAPIResponse
	err = json.Unmarshal(data, &apiResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse API response: %v", err)
	}
	result := WeatherResponse{
		City:        apiResponse.Name,
		Temperature: apiResponse.Main.Temp,
		Description: apiResponse.Weather[0].Description,
	}
	return &result, nil

} // why weatherresponse as pointer ? ans: We return a pointer to WeatherResponse to avoid copying the entire struct when returning from the function. This is more efficient, especially if the struct is large. Additionally, it allows us to return nil in case of an error, which can be a clear indication that something went wrong while fetching the weather data.
