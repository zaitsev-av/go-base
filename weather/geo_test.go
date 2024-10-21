package weather_test

import (
	"go-base/weather"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	//Arange - подготовка теста, предварительный результат (данные для функции)
	city := "London"
	expected := weather.GeoData{City: "London"}
	//Act - выполнение
	got, err := weather.GetMyLocation(city)
	if err != nil {
		t.Error("Ошибка получения города")
	}

	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получено %v", expected.City, got.City)
	}
	//Assert - на сколько совпадает с предварительным результатом
}
