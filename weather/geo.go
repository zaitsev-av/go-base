package weather

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

type CheckCity struct {
	Error bool `json:"error"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		err := checkCity(city)

		if !err {
			panic("Такого города не существует")
		}

		return &GeoData{
			City: city,
		}, nil
	}

	resp, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var responseData GeoData
	err = json.Unmarshal(body, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}

func checkCity(city string) bool {
	marshal, err := json.Marshal(map[string]string{
		"city": city,
	})
	if err != nil {
		return false
	}

	data, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(marshal))
	defer data.Body.Close()
	body, err := io.ReadAll(data.Body)
	if err != nil {
		return false
	}

	var response CheckCity
	err = json.Unmarshal(body, &response)
	return !response.Error
}
