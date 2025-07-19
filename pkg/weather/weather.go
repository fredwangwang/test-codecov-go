package weather

func GetWeather(city string) string {
	weathers := []string{"sunny", "cloudy", "rainy", "stormy", "snowy", "windy"}
	return weathers[len(city)%len(weathers)]
}
