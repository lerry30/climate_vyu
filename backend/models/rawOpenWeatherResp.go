package models

// shared struct
type Main struct {
	Temp float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	Humidity int `json:"humidity"`
	Pressure int `json:"pressure"`
}

type Weather []struct {
	Description string `json:"description"`
	Icon string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg int `json:"deg"`
}

type Clouds struct {
	All int `json:"all"`
}

// weather forecast list item
type RawWeatherForecastThreeHourInterval struct {
	DT int64 `json:"dt"`
	Main Main `json:"main"`
	Weather Weather `json:"weather"`
	Clouds Clouds `json:"clouds"`
	Wind Wind `json:"wind"`
	Visibility int `json:"visibility"`
	Pop float64 `json:"pop"`
	Sys struct {
		Pod string `json:"pod"`
	} `json:"sys"`
	DT_Txt string `json:"dt_txt"`
}

// -----------

type RawCurrentWeatherData struct {
	Name string `json:"name"`
	Main Main `json:"main"`
	Weather Weather `json:"weather"`
	Sys struct {
		Country string `json:"country"`
		Sunrise int64 `json:"sunrise"`
		Sunset int64 `json:"sunset"`
	} `json:"sys"`
	Wind Wind `json:"wind"`
	Clouds Clouds `json:"clouds"`
	Visibility int `json:"visibility"`
	DT int64 `json:"dt"`
	Timezone int `json:"timezone"`
}

type RawForecastWeatherData struct {
	COD string `json:"cod"`
	CNT int `json:"cnt"`
	List []RawWeatherForecastThreeHourInterval `json:"list"`
	City struct {
		Name string `json:"name"`
		Country string `json:"country"`
		Timezone int `json:"timezone"`
		Sunrise int64 `json:"sunrise"` 
		Sunset int64 `json:"sunset"`
	} `json:"city"`
}