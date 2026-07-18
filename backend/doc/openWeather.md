# OpenWeather API

This application uses OpenWeather API that allows 1000 request per day for their free trier. Since it just a small and personal app, using free tier is good for now.

## Struct

- id
- apiKey - OpenWeather's API key, requires account in openweathermap.org
- baseUrl - OpenWeather's url along with the url path and version <code>/data/2.5</code>
- httpClient - is a pointer of an instance of <code>http.Client</code> with the property of Timeout. Timeout defines how long to wait for the response after the request has been sent using the <code>Get(url)</code> method to fetch request from an API.

So, <code>http.Client</code> with <code>Get(url)</code> method is a way to send http request to an external API.

```go
type OpenWeather struct {
	id         string
	apiKey     string
	baseUrl    string
	httpClient *http.Client
}
```

## Constructor

Assigning all necessary values for <code>OpenWeather</code> properties.

```go
func NewOpenWeather(key string) *OpenWeather {
	return &OpenWeather{
		apiKey:  key,
		baseUrl: "https://api.openweathermap.org/data/2.5",
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
```

## <code>OpenWeather</code> Methods

<code>GetCurrent</code> method accept string, which in this case the city to search for weather forecast. It uses <code>fmt.Sprintf</code> to build the API URL with placeholder for required parameters.

The <code>ow.httpClient.Get(url)</code>(instance of <code>http.Client</code>) is the thing that sends the request. It has two return values, the first one is the http response and the second one is the error from that action.

<code>defer resp.Body.Close()</code> ensures that the body is closed regardless of how the method exits, even if the error occurs later.

Why the response <code>Body</code> should be closed when the operations are finished. <code>Close()</code> method will release the network resources and allowing the underlying TCP connection to be reused. <code>http.Client.Get(url)</code> is not actually finished yet on getting the entire response, think of it this way, if the response is 50MB, Go gives you stream <code>resp.Body</code>, instead of loading all at once, it gives you small chuck from the network. So, there are file descriptors(sockets) sitting around open, and this idle connections needs to be closed manually because if not, you'll run out of available sockets. Go uses pooling mechanism for handling HTTP responses, it doesn't necessarily read the entire response immedietely.

```go
func (ow *OpenWeather) GetCurrent(city string) (models.CurrentWeatherData, error) {
	city = url.QueryEscape(city)
	url := fmt.Sprintf("%s/weather?q=%s&appid=%s&units=metric", ow.baseUrl, city, ow.apiKey)

	resp, err := ow.httpClient.Get(url)
	if err != nil {
		return models.CurrentWeatherData{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.CurrentWeatherData{}, fmt.Errorf("openweather error: %s", resp.Status)
	}

	var data models.RawCurrentWeatherData

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return models.CurrentWeatherData{}, err
	}

	if len(data.Weather) == 0 {
		return models.CurrentWeatherData{}, fmt.Errorf("No weather data.")
	}

	wd := models.NewCurrentWeatherData()
	wd.TransformCurrentWeatherValues(&data)

	return *wd, nil
}
```


```go
func (ow *OpenWeather) GetForecast(city string) (models.ForecastWeatherData, error) {
	city = url.QueryEscape(city)
	url := fmt.Sprintf("%s/forecast?q=%s&appid=%s&units=metric", ow.baseUrl, city, ow.apiKey)

	resp, err := ow.httpClient.Get(url)
	if err != nil {
		return models.ForecastWeatherData{}, fmt.Errorf("openweather error: %s. City: %s", err, city)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.ForecastWeatherData{}, fmt.Errorf("failed to get weather forecast for %s.", city)
	}

	var data models.RawForecastWeatherData

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return models.ForecastWeatherData{}, fmt.Errorf("No weather data.")
	}

	if data.COD != "200" || len(data.List) == 0 {
		return models.ForecastWeatherData{}, fmt.Errorf("No weather data.")
	}

	fw := models.NewForecastWeatherData()
	fw.TransformForecastWeatherValues(&data)

	return *fw, nil
}
```