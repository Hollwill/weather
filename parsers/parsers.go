package parsers

import (
	"encoding/json"
	"github.com/HollWill/weather/structures"
)

type Parser interface {
	Parse(string) structures.WeatherData
	toWeatherData(interface{}) structures.WeatherData
}

type WeatherApiParser struct {
}

func (p WeatherApiParser) Parse(jsonString string) structures.WeatherData {

	dataWrapper := WeatherApiDataWrapper{}
	json.Unmarshal([]byte(jsonString), &dataWrapper)

	return p.toWeatherData(dataWrapper)
}

func (p WeatherApiParser) toWeatherData(dataWrapper interface{}) structures.WeatherData {
	wrapper := dataWrapper.(WeatherApiDataWrapper)
	data := wrapper.Data
	return structures.WeatherData{
		Temp:      data.Temp,
		FeelsLike: data.FeelsLike,
		Wind:      data.Wind,
		Text:      data.Condition.Text,
	}

}
