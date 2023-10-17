package models

import (
	"reflect"
	"testing"
)

func TestCalculateDeliveryTime(t *testing.T) {
	t.Run("no existing loads", func(t *testing.T) {
		load := Load{
			LoadId:  12,
			Pickup:  Location{Latitude: 50.0, Longitude: 50.0},
			Dropoff: Location{Latitude: 100.0, Longitude: 100.0}}

		assignment := &DriverDeliveryAssignment{DriverId: 1, Loads: []Load{}}
		got := roundTo2DecimalPlaces(assignment.CalculateDeliveryTime(load))
		want := 282.84

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("has existing loads", func(t *testing.T) {
		load := Load{
			LoadId:  12,
			Pickup:  Location{Latitude: 20.0, Longitude: 20.0},
			Dropoff: Location{Latitude: 100.0, Longitude: 100.0}}

		assignment := &DriverDeliveryAssignment{DriverId: 1, Loads: []Load{load}}

		got := roundTo2DecimalPlaces(assignment.CalculateDeliveryTime(load))
		want := 367.70

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("no existing loads, pass empty load", func(t *testing.T) {
		assignment := &DriverDeliveryAssignment{DriverId: 1, Loads: []Load{}}
		got := roundTo2DecimalPlaces(assignment.CalculateDeliveryTime(Load{}))
		want := 0.00

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

}

func TestGetDriverLocation(t *testing.T) {
	t.Run("no existing loads", func(t *testing.T) {

		assignment := &DriverDeliveryAssignment{DriverId: 1, Loads: []Load{}}
		got := assignment.GetDriverLocation()
		want := Location{0, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf(
				"should return depot location when there are no loads.\ngot %.2f want %.2f",
				got, want)
		}
	})

	t.Run("has existing loads", func(t *testing.T) {
		load := Load{
			LoadId:  12,
			Pickup:  Location{Latitude: 20.0, Longitude: 20.0},
			Dropoff: Location{Latitude: 100.0, Longitude: 100.0}}

		assignment := &DriverDeliveryAssignment{DriverId: 1, Loads: []Load{load}}

		got := assignment.GetDriverLocation()
		want := Location{Latitude: 100.0, Longitude: 100.0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf(
				"should return the dropoff point of the most recent load delivered.\n"+
					"got %.2f want %.2f", got, want)
		}
	})
}
