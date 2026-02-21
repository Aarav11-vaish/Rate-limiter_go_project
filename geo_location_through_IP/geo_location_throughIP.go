package geolocationthroughipgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type GeoLocation struct {
	Country string `json:"country"`
	Region  string `json:"regionName"`
}

func Getlocation(ip string) (string, error) {
	client := &http.Client{
		// fetch is a built in function in nodejs but in golang we have to create a client to make http requests and set the timeout for the request

		Timeout: time.Second * 4,
	}
	resp, err := client.Get("http://ip-api.com/json/" + ip)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close() // close the response body after we are done with it
	var data GeoLocation
	fmt.Println("Response : ", resp)
	// lets unmarshal the response body into a struct
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}
	return data.Country + " " + data.Region, err

}

// compare the above code with nodejs code->
// const response = await fetch(`https://ipapi.co/${ip}/json/`, { timeout: 4000 }); here we didn;t create any client because fetch is a built in function in nodejs but in golang we have to create a client to make http requests and set the timeout for the request
