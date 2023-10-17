package utils

import (
	. "com.vorto.vehiclerouter/pkg/models"
	"math"
	"reflect"
	"testing"
)

func TestTimeToDeliverLoad(t *testing.T) {
	start := Location{}
	load := Load{
		LoadId:  12,
		Pickup:  Location{Latitude: 50.0, Longitude: 50.0},
		Dropoff: Location{Latitude: 100.0, Longitude: 100.0}}

	got := roundTo2DecimalPlaces(TimeToDeliverLoad(start, load))
	want := 282.84

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestTotalCost(t *testing.T) {
	got := roundTo2DecimalPlaces(TotalCost(getMockDriverDeliveryAssignments()))
	want := 1032.44

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestDeliveryScheduleWithLoadIds(t *testing.T) {
	got := DeliveryScheduleWithLoadIds(getMockDriverDeliveryAssignments())
	want := [][]int{{1, 2, 4}, {3}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func roundTo2DecimalPlaces(number float64) float64 {
	return math.Round(number*100) / 100
}

func getMockDriverDeliveryAssignments() []DriverDeliveryAssignment {
	return []DriverDeliveryAssignment{
		{DriverId: 1, Loads: []Load{
			{LoadId: 1, Pickup: Location{-50.1, 80.0}, Dropoff: Location{90.1, 12.2}},
			{LoadId: 2, Pickup: Location{-24.5, -19.2}, Dropoff: Location{98.5, 1.8}},
			{LoadId: 4, Pickup: Location{0.3, 8.9}, Dropoff: Location{40.9, 55.0}}}, TotalDeliveryTime: 12.22},
		{DriverId: 2, Loads: []Load{{LoadId: 3, Pickup: Location{5.3, -61.1}, Dropoff: Location{77.8, -5.4}}}, TotalDeliveryTime: 20.22},
	}
}
