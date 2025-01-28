package main

import (
	"WeatherAPI/internal/pkg/WeatherApi"
)

func main() {

	app := WeatherApi.New()

	app.Run()

}
