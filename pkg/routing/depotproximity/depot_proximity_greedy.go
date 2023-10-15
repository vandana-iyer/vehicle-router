package depotproximity

import (
	. "com.vorto.vehiclerouter/pkg/models"
	"com.vorto.vehiclerouter/pkg/routing"
	"com.vorto.vehiclerouter/pkg/utils"
	"sort"
)

// DepotProximityGreedy is a greedy routing strategy that prioritizes loads
// based on their proximity from the depot to the pickup location.
type DepotProximityGreedy struct{}

func (bg *DepotProximityGreedy) ScheduleLoads(loads []Load) []DriverDeliveryAssignment {
	var assignments []DriverDeliveryAssignment
	assignment := DriverDeliveryAssignment{DriverId: 1}
	currentLocation := Location{}

	// Sort the loads based on distance from depot to pickup location
	sort.Slice(loads, func(i, j int) bool {
		return currentLocation.TravelDuration(loads[i].Pickup) > currentLocation.TravelDuration(loads[j].Pickup)
	})

	for _, load := range loads {
		deliveryTimeFromCurrentLocation := utils.TimeToDeliverLoad(currentLocation, load)

		if assignment.TotalDeliveryTime+deliveryTimeFromCurrentLocation > routing.MaxTime {
			assignments = append(assignments, assignment)
			assignment = DriverDeliveryAssignment{DriverId: assignment.DriverId + 1}
			currentLocation = Location{}
		}

		assignment.Loads = append(assignment.Loads, load)
		assignment.TotalDeliveryTime += deliveryTimeFromCurrentLocation
		currentLocation = load.Dropoff
	}

	if len(assignment.Loads) > 0 {
		assignments = append(assignments, assignment)
	}
	return assignments
}
