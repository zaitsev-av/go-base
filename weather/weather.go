package weather

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Weather() {
	city := flag.String("city", "", "Город")
	format := flag.Int("format", 4, "Формат вывода погоды")
	flag.Parse()

	location, _ := GetMyLocation(*city)
	getWeather(*location, *format)
}

func getWeather(geo GeoData, format int) string {
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		fmt.Println("Не удалось спарсить URL")
		return ""
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err)
		return ""
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println(string(body))

	return string(body)
}
