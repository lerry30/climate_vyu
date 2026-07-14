package models

import (
	"fmt"
	"time"

	"backend/utils"
)

type CurrentWeatherData struct {
	CityName string `json:"city_name"`
	Temp string `json:"temp"`
	FeelsLike string `json:"feels_like"`
	Description string `json:"description"`
	Humidity string `json:"humidity"`
	Pressure string `json:"pressure"`
	Icon string `json:"icon"`
	Country string `json:"country"`
	Sunrise time.Time `json:"sunrise"`
	Sunset time.Time `json:"sunset"`
	WindSpeed string `json:"wind_speed"`
	WindDirection string `json:"wind_direction"`
	CloudsCover string `json:"clouds_cover"`
	Visibility string `json:"visibility"`
	LastUpdateTime time.Time `json:"last_update_time"`
}

func NewCurrentWeatherData() *CurrentWeatherData {
	return &CurrentWeatherData{}
}

func (wd *CurrentWeatherData) TransformCurrentWeatherValues(raw *RawCurrentWeatherData) {
	if len(raw.Weather) == 0 {
		return
	}

	//const kDiff = 273.15 // C = K - 273.15
	temp := fmt.Sprintf("%.2f°C", raw.Main.Temp)
	feelsLike := fmt.Sprintf("%.2f°C", raw.Main.FeelsLike)
	humidity := fmt.Sprintf("%d%%", raw.Main.Humidity)
	pressure := fmt.Sprintf("%d%% hPa", raw.Main.Pressure)

	icon := fmt.Sprintf("http://openweathermap.org/img/w/%s", raw.Weather[0].Icon)

	sunrise := utils.ToLocalTime(raw.Sys.Sunrise, raw.Timezone)
	sunset := utils.ToLocalTime(raw.Sys.Sunset, raw.Timezone)
	
	speed := fmt.Sprintf("%.2f m/s", raw.Wind.Speed)
	windDirection := fmt.Sprintf("%d°", raw.Wind.Deg)
	
	clouds := fmt.Sprintf("%d%%", raw.Clouds.All)
	visibility := fmt.Sprintf("%.2f km", float64(raw.Visibility) / 1000.0)
	
	lastUpdateTime := utils.ToLocalTime(raw.DT, raw.Timezone)

	wd.CityName = raw.Name
	wd.Temp = temp
	wd.FeelsLike = feelsLike
	wd.Description = raw.Weather[0].Description
	wd.Humidity = humidity
	wd.Pressure = pressure
	wd.Icon = icon
	wd.Country = raw.Sys.Country
	wd.Sunrise = sunrise
	wd.Sunset = sunset
	wd.WindSpeed = speed
	wd.WindDirection = windDirection
	wd.CloudsCover = clouds
	wd.Visibility = visibility
	wd.LastUpdateTime = lastUpdateTime
}