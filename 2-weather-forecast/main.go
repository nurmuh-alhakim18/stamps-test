package main

import (
	"2-weather-forecast/config"
	"2-weather-forecast/external"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	client := external.NewWeatherClient(cfg.ApiKey)
	forecasts, err := client.GetForecasts("celcius")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Weather Forecasts:")
	processForecasts(forecasts)
}

func processForecasts(forecasts *external.Forecast) {
	for i := 0; i < len(forecasts.List); i++ {
		dateTime := forecasts.List[i].DtTxt
		if strings.Contains(dateTime, "12:00:00") {
			readableDate, err := convertDateTime(dateTime)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%s: %.2f°C\n", readableDate, forecasts.List[i].Main.Temp)
		}
	}
}

func convertDateTime(dateTimeStr string) (string, error) {
	time, err := time.Parse("2006-01-02 15:04:05", dateTimeStr)
	if err != nil {
		errMsg := fmt.Sprintf("failed to make parse date time: %v", err)
		return "", errors.New(errMsg)
	}

	weekday := time.Weekday()
	day := time.Day()
	month := time.Month()
	year := time.Year()
	return fmt.Sprintf("%s, %d %v %d", weekday, day, month, year), nil
}