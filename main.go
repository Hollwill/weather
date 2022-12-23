package main

import (
	"fmt"
	"log"
	"os"

	"github.com/HollWill/weather/api_services"
	"github.com/HollWill/weather/structures"
)

func main() {
	key, ok := os.LookupEnv("WEATHER_API_KEY")
	if !ok {
		log.Default().Fatalln("Add WEATHER_API_KEY to your environment variables")
	}

	city := os.Args[1]

	var t api_services.Service
	t = api_services.WeatherApiComService{Coords: structures.Coords{City: city}, ApiKey: key}
	jsonString := t.FetchData()

	parser := t.GetParser()

	parsedStructure := parser.Parse(jsonString)
	fmt.Println(parsedStructure)
}
