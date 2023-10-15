package loadproximity

import (
	. "com.vorto.vehiclerouter/pkg/models"
	"com.vorto.vehiclerouter/pkg/routing/depotproximity"
	"com.vorto.vehiclerouter/pkg/utils"
	"reflect"
	"testing"
)

func TestScheduleLoads(t *testing.T) {
	t.Run("collection of 4 loads", func(t *testing.T) {
		loads := []Load{
			{1, Location{-50.1, 80.0}, Location{90.1, 12.2}},
			{2, Location{-24.5, -19.2}, Location{98.5, 1.8}},
			{3, Location{0.3, 8.9}, Location{40.9, 55.0}},
			{4, Location{5.3, -61.1}, Location{77.8, -5.4}},
		}

		strategyInstance := ClosestLoadGreedy{}

		got := utils.DeliveryScheduleWithLoadIds(strategyInstance.ScheduleLoads(loads))
		want := [][]int{
			{3, 1, 4},
			{2},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v want %+v", got, want)
		}
	})

	t.Run("no loads", func(t *testing.T) {
		loads := []Load{}

		strategyInstance := depotproximity.DepotProximityGreedy{}

		got := strategyInstance.ScheduleLoads(loads)

		if got != nil {
			t.Errorf("got %+v, want a nil return", got)
		}
	})
}
