package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var apiKey, part string
var lat, long = 37.4419, -122.1430

const (
	IpInfoUrl     = "https://ipinfo.io/json"
	WeatherApiUrl = "https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&exclude=%s&appid=%s"
)

type IPInfoResponse struct {
	Loc string `json:"loc"`
}

type WeatherApiResponse struct {
	Weather []struct {
		Condition   string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func SetCoordinates(w http.ResponseWriter, r *http.Request) {
	// set up coordinates
	resp, err := http.Get(IpInfoUrl)
	if err != nil {
		fmt.Println("failed to fetch location", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("failed to read response body", err)
		return
	}

	var ipInfo IPInfoResponse
	if err := json.Unmarshal(body, &ipInfo); err != nil {
		fmt.Println("failed to parse JSON for info api", err)
		return
	}

	locArr := strings.Split(ipInfo.Loc, ",")
	lat, err = strconv.ParseFloat(locArr[0], 64)
	if err != nil {
		fmt.Println("failed to parse latitude", err)
		return
	}

	long, err = strconv.ParseFloat(locArr[1], 64)
	if err != nil {
		fmt.Println("failed to parse latitude", err)
		return
	}
}

func getTemperatureStatus(temp float64) string {
	// comparing temperatures
	if temp > 120 {
		return "hot"
	} else if temp < 60 {
		return "cold"
	} else {
		return "moderate"
	}
}

func WeatherApiHandler(w http.ResponseWriter, r *http.Request) {
	// api to give weather conditions
	apiKey := r.URL.Query().Get("apiKey")
	part := r.URL.Query().Get("part")

	url := fmt.Sprintf(WeatherApiUrl, lat, long, part, apiKey)
	fmt.Println("my url", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("failed to fetch location", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("failed to read response body", err)
		return
	}

	var weatherResponse WeatherApiResponse
	if err := json.Unmarshal(body, &weatherResponse); err != nil {
		fmt.Println("failed to parse JSON for weather api", err)
		return
	}

	tempStatus := getTemperatureStatus(weatherResponse.Main.Temp)

	response := fmt.Sprintf("Weather api shows your weather condition outside is %s, %s and temperature is %s \n",
		weatherResponse.Weather[0].Condition, weatherResponse.Weather[0].Description, tempStatus)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		SetCoordinates(w, r)
		WeatherApiHandler(w, r)
	})

	fmt.Println("starting server at 8080")
	http.ListenAndServe(":8080", mux)

}
