package weather

import (
	"math/rand"
	"time"
)

func GetWeather(city string) string {
	rand.Seed(time.Now().UnixNano())
	weathers := []string{"sunny", "cloudy", "rainy", "stormy", "snowy", "windy"}
	return weathers[rand.Intn(len(weathers))]
}
