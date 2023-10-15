package utils

import "com.vorto.vehiclerouter/pkg/models"

const CostPerDriver = 500.0

// TimeToDeliverLoad helps compute the total time taken to deliver a load
// from a given location upto a load's dropoff location.
func TimeToDeliverLoad(currentLocation models.Location, load models.Load) float64 {
	pickupDuration := currentLocation.TravelDuration(load.Pickup)
	dropoffDuration := load.Pickup.TravelDuration(load.Dropoff)
	return pickupDuration + dropoffDuration
}

func TotalCost(assignments []models.DriverDeliveryAssignment) float64 {
	driverCost := CostPerDriver * float64(len(assignments))

	totalDeliveryTime := 0.0
	for _, assignment := range assignments {
		totalDeliveryTime += assignment.TotalDeliveryTime
	}
	return driverCost + totalDeliveryTime
}

func DeliveryScheduleWithLoadIds(driversAssignment []models.DriverDeliveryAssignment) [][]int {
	var loadIDsList [][]int

	for _, driverAssignment := range driversAssignment {
		loadIDsList = append(loadIDsList, getLoadIds(driverAssignment.Loads))
	}
	return loadIDsList
}

func getLoadIds(loads []models.Load) []int {
	loadIds := make([]int, len(loads))
	for i, load := range loads {
		loadIds[i] = load.LoadId
	}
	return loadIds
}
