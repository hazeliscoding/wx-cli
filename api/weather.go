package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type WeatherResponse struct {
	Location struct {
		Name    string `json:"name"`
		Region  string `json:"region"`
		Country string `json:"country"`
		TzID    string `json:"tz_id"`
	} `json:"location"`
	Current struct {
		TempF      float64 `json:"temp_f"`
		FeelsLikeF float64 `json:"feelslike_f"`
		Humidity   int     `json:"humidity"`
		Condition  struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		ForecastDay []struct {
			Date string `json:"date"`
			Day  struct {
				AvgTempF      float64 `json:"avgtemp_f"`
				MinTempF      float64 `json:"mintemp_f"`
				MaxTempF      float64 `json:"maxtemp_f"`
				TotalPrecipIn float64 `json:"totalprecip_in"`
				Condition     struct {
					Text string `json:"text"`
				} `json:"condition"`
			} `json:"day"`
			Astro struct {
				Sunrise string `json:"sunrise"`
				Sunset  string `json:"sunset"`
			} `json:"astro"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func FetchWeather(apiKey, city, region string) (*WeatherResponse, error) {
	encodedCity := url.QueryEscape(city + ", " + region)
	url := fmt.Sprintf("https://api.weatherapi.com/v1/forecast.json?q=%s&days=7&key=%s", encodedCity, apiKey)
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("error fetching weather data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	var weather WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, fmt.Errorf("error decoding API response: %v", err)
	}

	return &weather, nil
}
