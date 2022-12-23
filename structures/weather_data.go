package structures

import "fmt"

type WeatherData struct {
	Temp      float32
	FeelsLike float32
	Wind      float32
	Text      string
}

func (d WeatherData) String() string {
	return fmt.Sprintf("Температура: %v°C %v \nЧувствуется как: %v°C \nВетер: %vм/с", d.Temp, d.Text, d.FeelsLike, d.Wind)
}
