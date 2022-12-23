package api_services

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/HollWill/weather/parsers"
	"github.com/HollWill/weather/structures"
)

type Service interface {
	FetchData() string
	GetParser() parsers.Parser
}

const WEATHER_API_URL string = "http://api.weatherapi.com/v1/current.json?lang=ru&key=%s&q=%v"

type WeatherApiComService struct {
	Coords structures.Coords
	ApiKey string
}

func (w WeatherApiComService) FetchData() string {
	return w.makeRequest(w.Coords)
}

func (w WeatherApiComService) makeRequest(coords structures.Coords) string {
	searchString := w.getSearchString()

	response, err := http.Get(fmt.Sprintf(WEATHER_API_URL, w.ApiKey, searchString))
	if err != nil {
		log.Fatalln(err)
	}

	return w.getStringFromResponse(*response)
}

func (w WeatherApiComService) getSearchString() string {
	if w.Coords.City != "" {
		return w.Coords.City
	} else {
		return fmt.Sprintf("%v,%v", w.Coords.Lat, w.Coords.Long)
	}
}

func (w WeatherApiComService) getStringFromResponse(resp http.Response) string {

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}

func (w WeatherApiComService) GetParser() parsers.Parser {
	return parsers.WeatherApiParser{}
}
