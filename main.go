package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type WeatherResponse struct {
	Name string `json:"name"`
	Weather Weather `json:"weather"`
	Main Main	`json:"main"`
}

type Weather []struct {
	Description string `json:"description"`
}

type Main struct {
	Temp float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin float64 `json:"temp_min"`
	TempMax float64 `json:"temp_max"`
	Pressure float64 `json:"pressure"`
	Humidity float64 `json:"humidity"`
}

func main() {
	var city string
	fmt.Print("Enter city name: ")
	fmt.Scanln(&city)

	results := getWeatherData(city)
	fmt.Println("Weather: ", results.Weather[0].Description)
	fmt.Println("Temperature: ", results.Main.Temp-273.15)
}

func getWeatherData(city string) WeatherResponse {
	API_KEY := "f1186383ba40084e342ccc135b8aa771"
    BASE_URL := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?appid=%s&q=%s", API_KEY, city)
	response, err := http.Get(BASE_URL)
	if err != nil {
		log.Fatal(err)
	}

	jsonBytes, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	var weatherResponse WeatherResponse
	er := json.Unmarshal(jsonBytes, &weatherResponse)
	if er != nil {
		log.Fatal(er)
	}

	return weatherResponse
}