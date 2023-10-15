package models

// DriverDeliveryAssignment represents the loads scheduled for a particular driver,
// and the total time it will take to deliver all those loads.
type DriverDeliveryAssignment struct {
	DriverId          int
	Loads             []Load
	TotalDeliveryTime float64
}
