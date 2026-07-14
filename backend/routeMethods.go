package main

import (
	"backend/models"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

// Gives suggestions for the available cities
func (apiServer *APIServer) FilterCity(w http.ResponseWriter, r *http.Request) error {
	// this struct city holds the city name from the request URL parameter
	var city models.City

	// mux to get the necessary parameters
	vars := mux.Vars(r)
	cityName := vars["city"]

	// remove white spaces on both ends
	city.Name = strings.TrimSpace(cityName)

	if city.Name == "" {
		fmt.Fprintf(os.Stderr, "Error: empty string")
		//log.Println("Error: empty string")
		return fmt.Errorf("Empty data.")
	}

	// this struct holds the city names for suggestions
	var fData models.FilteredCities
	fData, err := city.Filter() // return: struct, error
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %w", err.Error())
		//log.Println("=======Error", err.Error(), "========")
		return fmt.Errorf("City name not found.")
	}

	// response
	ResponseJSON(w, http.StatusOK, fData)

	return nil
}

func (apiServer *APIServer) SearchCity(w http.ResponseWriter, r *http.Request) error {
	// mux to get the request URL parameters needed, and in this case, the city name
	vars := mux.Vars(r)
	city := vars["city"]

	// remove white spaces on both ends
	city = strings.TrimSpace(city)

	if city == "" {
		fmt.Fprintf(os.Stderr, "Empty string: city")
		return fmt.Errorf("Empty data.")
	}

	// get the external API struct pointer, which is then configured in the main function and added in api server
	// the ok return parameter provides a boolean status of getting this struct from this map of type any, whether the struct exists or not
	externalApi, ok := apiServer.externals["open-weather"]
	if !ok || externalApi == nil {
		fmt.Fprintf(os.Stderr, "open-weather is not exists in external apis")
		return fmt.Errorf("Internal server error")
	}

	// the ok holds the boolean status of converting the type any to a pointer of some sort of type struct
	ow, ok := externalApi.(*OpenWeather)
	if !ok {
		fmt.Fprintf(os.Stderr, "failed to convert any to *OpenWeather")
		return fmt.Errorf("Weather service not available.")
	}

	// wd has all the necessary fields needed to be filled by GetCurrent
	var wd models.CurrentWeatherData
	// ow is a struct for external API requests
	// GetCurrent gives the weather update for a specific city
	wd, err := ow.GetCurrent(city)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %w", err)
		return fmt.Errorf("No weather data.")
	}

	// response
	ResponseJSON(w, http.StatusOK, wd)

	return nil
}

func (apiServer *APIServer) Forecast(w http.ResponseWriter, r *http.Request) error {
	// mux to get the request URL parameters needed, and in this case, the city name
	vars := mux.Vars(r)
	city := vars["city"]

	// remove white spaces on both ends
	city = strings.TrimSpace(city)

	if city == "" {
		fmt.Fprintf(os.Stderr, "Empty string: city")
		return fmt.Errorf("Empty data.")
	}

	// get the external API struct pointer, which is then configured in the main function and added in api server
	// the ok return parameter provides a boolean status of getting this struct from this map of type any, whether the struct exists or not
	externalApi, ok := apiServer.externals["open-weather"]
	if !ok || externalApi == nil {
		fmt.Fprintf(os.Stderr, "open-weather is not exists in external apis")
		return fmt.Errorf("Internal server error")
	}

	// the ok holds the boolean status of converting the type any to a pointer of some sort of type struct
	ow, ok := externalApi.(*OpenWeather)
	if !ok {
		fmt.Fprintf(os.Stderr, "failed to convert any to *OpenWeather")
		return fmt.Errorf("Weather service not available.")
	}

	// fw has all the necessary fields needed to be filled by GetForecast
	var fw models.ForecastWeatherData
	// ow is a struct for external API requests
	// GetForecast gives the weather forecast update for a specific city
	fw, err := ow.GetForecast(city)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %w", err)
		return fmt.Errorf("No weather data.")
	}

	// response
	ResponseJSON(w, http.StatusOK, fw)

	return nil
}
