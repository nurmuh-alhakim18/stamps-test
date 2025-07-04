package external

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type WeatherClient struct {
	httpClient http.Client
	apiKey 		 string
	unitType   map[string]string
}

func NewWeatherClient(apiKey string) *WeatherClient {
	return &WeatherClient{
		httpClient: http.Client{
			Timeout: 10 * time.Second,
		},
		apiKey: apiKey,
		unitType: UnitType(),
	}
}

func UnitType() map[string]string {
	return  map[string]string{
		"kelvin": "standard",
		"celcius": "metric",
		"fahrenheit": "imperial",
	}
}

func (wc *WeatherClient) GetForecasts(temperatureUnit string) (*Forecast, error) {
	cityID := "1642911"
	unit, ok := wc.unitType[temperatureUnit]
	if !ok {
		return nil, errors.New("unrecognized unit")
	}
	
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?id=%s&units=%s&appid=%s", cityID, unit, wc.apiKey)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		errMsg := fmt.Sprintf("failed to make request: %v", err)
		return nil, errors.New(errMsg)
	}

	resp, err := wc.httpClient.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("failed to send request: %v", err)
		return nil, errors.New(errMsg)
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		errMsg := fmt.Sprintf("error with status code: %d", resp.StatusCode)
		return nil, errors.New(errMsg)
	}

	var data Forecast
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, errors.New("failed to decode json")
	}

	return &data, nil
}