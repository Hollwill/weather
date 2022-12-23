package parsers

type WeatherApiDataWrapper struct {
	Data WeatherApiData `json:"current"`
}

type WeatherApiData struct {
	Temp      float32 `json:"temp_c"`
	FeelsLike float32 `json:"feelslike_c"`
	Wind      float32 `json:"wind_kph"`
	IsDay     bool    `json:"is_day"`
	Condition struct {
		Text string `json:"text"`
	} `json:"condition"`
}
