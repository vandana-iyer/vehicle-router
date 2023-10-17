package models

// DriverDeliveryAssignment represents the loads scheduled for a particular driver,
// and the total time it will take to deliver all those loads.
type DriverDeliveryAssignment struct {
	DriverId          int
	Loads             []Load
	TotalDeliveryTime float64
	Depot             Depot // location from which a driver starts and concludes their shift.
}

// CalculateDeliveryTime calculates the total duration required to transport a load
// from a specified starting location to the load's drop-off point and then
// return to the Depot.
func (assignment *DriverDeliveryAssignment) CalculateDeliveryTime(load Load) float64 {
	currentLocation := assignment.GetDriverLocation()
	pickupDuration := currentLocation.TravelDuration(load.Pickup)
	dropoffDuration := load.Pickup.TravelDuration(load.Dropoff)
	depotReturnDuration := load.Dropoff.TravelDuration(assignment.Depot.Location)

	return pickupDuration + dropoffDuration + depotReturnDuration
}

// getCurrentLocation retrieves the driver's current location, typically
// the most recent drop-off point. If the driver hasn't commenced their shift,
// the location defaults to the Depot location.
func (assignment *DriverDeliveryAssignment) GetDriverLocation() Location {
	var currentLocation Location
	if len(assignment.Loads) == 0 {
		currentLocation = assignment.Depot.Location
	} else {
		// last dropoff location
		currentLocation = assignment.Loads[len(assignment.Loads)-1].Dropoff
	}
	return currentLocation
}
