package models

import (
	"math"
	"testing"
)

func TestDuration(t *testing.T) {
	location := Location{50, 50}
	got := roundTo2DecimalPlaces(location.TravelDuration(Location{100, 100}))
	want := 70.71

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}

}

func roundTo2DecimalPlaces(number float64) float64 {
	return math.Round(number*100) / 100
}
