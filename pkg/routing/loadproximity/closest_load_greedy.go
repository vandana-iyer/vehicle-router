package loadproximity

import (
	. "com.vorto.vehiclerouter/pkg/models"
	"com.vorto.vehiclerouter/pkg/routing"
)

// ClosestLoadGreedy is a greedy routing strategy that continuously selects
// the load closest to the current location of the driver. This strategy is
// particularly effective for dense delivery zones where loads are in close
// distance to each other.
type ClosestLoadGreedy struct{}

func (cg *ClosestLoadGreedy) ScheduleLoads(loads []Load) []*DriverDeliveryAssignment {
	var assignments []*DriverDeliveryAssignment
	assignment := &DriverDeliveryAssignment{DriverId: 1}

	for len(loads) > 0 {
		// Find the job closest to the current location
		closestLoadIndex := findClosestLoadIndex(assignment.GetDriverLocation(), loads)
		load := loads[closestLoadIndex]

		deliveryTimeFromCurrentLocation := assignment.CalculateDeliveryTime(load)

		if assignment.TotalDeliveryTime+deliveryTimeFromCurrentLocation > routing.MaxTime {
			assignments = append(assignments, assignment)
			assignment = &DriverDeliveryAssignment{DriverId: assignment.DriverId + 1}
			continue
		}

		assignment.Loads = append(assignment.Loads, load)
		assignment.TotalDeliveryTime += deliveryTimeFromCurrentLocation

		// Remove the load from the list of pending loads
		loads = append(loads[:closestLoadIndex], loads[closestLoadIndex+1:]...)
	}

	if len(assignment.Loads) > 0 {
		assignments = append(assignments, assignment)
	}
	return assignments
}

// This function returns the index of the load closest to the given location
func findClosestLoadIndex(location Location, loads []Load) int {
	minDistance := float64(1<<63 - 1) // Start with max float value
	closestLoadIndex := 0

	for i, load := range loads {
		distance := location.TravelDuration(load.Pickup)
		if distance < minDistance {
			minDistance = distance
			closestLoadIndex = i
		}
	}

	return closestLoadIndex
}
