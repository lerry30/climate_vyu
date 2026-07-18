# Forecast Weather

What we got from Open Weather API's weather forecast data:

We have 40 items in a list, and this list represents 5 days of the weather forecast. 
So, 40/5 is 8, 8 sets of data in a day, which means that we have reports every 3 hours in a day.

## Structs

A struct for 3 hour weather data. 

```go
type WeatherForecastThreeHourInterval struct {
	Temp               string    `json:"temp"`
	FeelsLike          string    `json:"feels_like"`
	Humidity           string    `json:"humidity"`
	Pressure           string    `json:"pressure"`
	Description        string    `json:"description"`
	Icon               string    `json:"icon"`
	CloudsCover        string    `json:"clouds_cover"`
	WindSpeed          string    `json:"wind_speed"`
	WindDirection      string    `json:"wind_direction"`
	Visibility         string    `json:"visibility"`
	ChanceOfRainOrSnow string    `json:"chance_of_rain_or_snow"`
	PartOfDay          string    `json:"part_of_day"`
	DateTimeGMT        string    `json:"date_time"`
	LocalDateTime      time.Time `json:"local_date_time"`
}
```

Day of the week simply implies one of the seven named days(Monday-Sunday).
The <code>DayOfTheWeek</code> is used as a key to easily map and display specific day.
<code>HourWeatherUpdates</code> contains 8 sets of weather data in a day.

```go
type WeatherForecastWeekDay struct {
	DayOfTheWeek       string                             `json:"day_of_the_week"`
	HourWeatherUpdates []WeatherForecastThreeHourInterval `json:"hour_weather_updates"`
}
```

Main struct with all the properties needed.

```go
type ForecastWeatherData struct {
	CityName string                   `json:"city_name"`
	Country  string                   `json:"country"`
	Sunrise  time.Time                `json:"sunrise"`
	Sunset   time.Time                `json:"sunset"`
	List     []WeatherForecastWeekDay `json:"list"`
}
```

## Constructor

```go
func NewForecastWeatherData() *ForecastWeatherData {
	return &ForecastWeatherData{}
}
```

## Method

<code>TransformForecastWeatherValues(*RawForecastWeatherData)</code> transforms raw weather data to human readable data. But the main goal here is to structured the data so it can easily display in the dashboard.

```go
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
		// Day of the week simply implies one of the seven named days(Monday-Sunday).
		// So, day of the week could be Sunday or any of the days in a week.
		// The key variable holds this value.

		// local datetime
		dt := utils.ToLocalTime(item.DT, timezone)
		key := dt.Weekday().String()

		var dayOfTheWeek *WeatherForecastWeekDay // pointer, initialize nil by default

		// search through the list if the key exists
		// if it does, the pointer for that will be assigned to dayOfTheWeek
		for i, wd := range fw.List {
			if wd.DayOfTheWeek == key {
				dayOfTheWeek = &fw.List[i]
				break
			}
		}

		// if the key is not found in the list, this block will initialize a new
		// instance of WeatherForecastWeekDay for that key and its values.
		if dayOfTheWeek == nil {
			newWeekDay := WeatherForecastWeekDay{
				DayOfTheWeek:       key,
				HourWeatherUpdates: []WeatherForecastThreeHourInterval{},
			}

			// assign an empty struct
			fw.List = append(fw.List, newWeekDay)
			dayOfTheWeek = &newWeekDay
		}

		// If there's no weather data for 3 hours then it will assigned and empty struct and skip the bottom part.
		if len(item.Weather) == 0 {
			dayOfTheWeek.HourWeatherUpdates = append(dayOfTheWeek.HourWeatherUpdates, WeatherForecastThreeHourInterval{})
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

		visibility := fmt.Sprintf("%.2f km", float64(item.Visibility)/1000.0)

		pop := fmt.Sprintf("%.2f%%", item.Pop)

		pod := "day"
		if item.Sys.Pod == "n" {
			pod = "night"
		}
		// ---

		// append a new item to *WeatherForecastWeekDay
		dayOfTheWeek.HourWeatherUpdates = append(dayOfTheWeek.HourWeatherUpdates, WeatherForecastThreeHourInterval{
			Temp:               temp,
			FeelsLike:          feelsLike,
			Humidity:           humidity,
			Pressure:           pressure,
			Description:        item.Weather[0].Description,
			Icon:               icon,
			CloudsCover:        clouds,
			WindSpeed:          speed,
			WindDirection:      windDirection,
			Visibility:         visibility,
			ChanceOfRainOrSnow: pop,
			PartOfDay:          pod,
			DateTimeGMT:        item.DT_Txt, // GMT: ref time
			LocalDateTime:      dt,
		})
	}

	//d, _ := json.MarshalIndent(fw, "", "	")
	//fmt.Println(string(d))
}
```