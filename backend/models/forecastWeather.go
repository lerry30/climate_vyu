package models

import (
	"fmt"
	"time"

	"backend/utils"
)

// We have 40 items in a list, and this list represents 5 days of the weather forecast. 
// So, 40/5 is 8, 8 sets of data in a day, which means that we have a report every 3 hours 
// in a day.

type WeatherForecastThreeHourInterval struct {
	Temp string `json:"temp"`
	FeelsLike string `json:"feels_like"`
	Humidity string `json:"humidity"`
	Pressure string `json:"pressure"`
	Description string `json:"description"`
	Icon string `json:"icon"`
	CloudsCover string `json:"clouds_cover"`
	WindSpeed string `json:"wind_speed"`
	WindDirection string `json:"wind_direction"`
	Visibility string `json:"visibility"`
	ChanceOfRainOrSnow string `json:"chance_of_rain_or_snow"`
	PartOfDay string `json:"part_of_day"`
	DateTimeGMT string `json:"date_time"`
	LocalDateTime time.Time `json:"local_date_time"`
}

type WeatherForecastWeekDay struct {
	WeekDay string `json:"weekday"`
	HourWeatherUpdates []WeatherForecastThreeHourInterval `json:"hour_weather_updates"`
}

type ForecastWeatherData struct {
	CityName string `json:"city_name"`
	Country string `json:"country"`
	Sunrise time.Time `json:"sunrise"`
	Sunset time.Time `json:"sunset"`
	List []WeatherForecastWeekDay `json:"list"`
}

func NewForecastWeatherData() *ForecastWeatherData {
	return &ForecastWeatherData{}
}

func (fw *ForecastWeatherData) TransformForecastWeatherValues(raw *RawForecastWeatherData) {
	if len(raw.List) == 0 {
		return
	}

	var timezone int = raw.City.Timezone
	sunrise := utils.ToLocalTime(raw.City.Sunrise, timezone)
	sunset := utils.ToLocalTime(raw.City.Sunset, timezone)
	
	fw.CityName = raw.City.Name
	fw.Country = raw.City.Country
	fw.Sunrise = sunrise
	fw.Sunset = sunset

	fw.List = []WeatherForecastWeekDay{}

	for _, item := range raw.List {
		// local datetime
		dt := utils.ToLocalTime(item.DT, timezone)
		key := dt.Weekday().String()

		var weekDay *WeatherForecastWeekDay // pointer, initialize nil by default

		// search through the list if the key exists
		// if it does, the pointer for that will be assigned to weekDay
		for i, wd := range fw.List {
			if wd.WeekDay == key {
				weekDay = &fw.List[i]
				break
			}
		}

		// if the key is not found in the list, this block will initialize a new
		// instance of WeatherForecastWeekDay for that key and its values.
		if weekDay == nil {
			newWeekDay := WeatherForecastWeekDay{
				WeekDay: key,
				HourWeatherUpdates: []WeatherForecastThreeHourInterval{},
			}

			// assign an empty struct
			fw.List = append(fw.List, newWeekDay)
			weekDay = &newWeekDay
		}

		if len(item.Weather) == 0 {
			weekDay.HourWeatherUpdates = append(weekDay.HourWeatherUpdates, WeatherForecastThreeHourInterval{})
			continue
		}

		// ---
		temp := fmt.Sprintf("%.2f°C", item.Main.Temp)
		feelsLike := fmt.Sprintf("%.2f°C", item.Main.FeelsLike)
		humidity := fmt.Sprintf("%d%%", item.Main.Humidity)
		pressure := fmt.Sprintf("%d%% hPa", item.Main.Pressure)

		icon := fmt.Sprintf("https://openweathermap.org/img/w/%s.png", item.Weather[0].Icon)

		clouds := fmt.Sprintf("%d%%", item.Clouds.All)
		
		speed := fmt.Sprintf("%.2f m/s", item.Wind.Speed)
		windDirection := fmt.Sprintf("%d°", item.Wind.Deg)
		
		visibility := fmt.Sprintf("%.2f km", float64(item.Visibility) / 1000.0)

		pop := fmt.Sprintf("%.2f%%", item.Pop)

		pod := "day"
		if item.Sys.Pod == "n" {
			pod = "night"
		}
		// ---

		// append a new item to *WeatherForecastWeekDay
		weekDay.HourWeatherUpdates = append(weekDay.HourWeatherUpdates, WeatherForecastThreeHourInterval{
			Temp: temp,
			FeelsLike: feelsLike,
			Humidity: humidity,
			Pressure: pressure,
			Description: item.Weather[0].Description,
			Icon: icon,
			CloudsCover: clouds,
			WindSpeed: speed,
			WindDirection: windDirection,
			Visibility: visibility,
			ChanceOfRainOrSnow: pop,
			PartOfDay: pod,
			DateTimeGMT: item.DT_Txt, // GMT: ref time
			LocalDateTime: dt,
		})
	}
}