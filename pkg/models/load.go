package models

// Load represents an item that needs to be transported from a pickup location
// to a dropoff location. Each load has a unique identifier, LoadId.
type Load struct {
	LoadId  int
	Pickup  Location
	Dropoff Location
}
