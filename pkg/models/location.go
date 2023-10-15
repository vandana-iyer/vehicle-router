package models

import "math"

// Location represents a geographical point defined by its latitude and longitude coordinates.
type Location struct {
	Latitude  float64
	Longitude float64
}

// TravelDuration returns time taken to travel from one location to another
// by calculating the Euclidean distance between both the locations.
func (currentLocation Location) TravelDuration(destination Location) float64 {
	travelDuration := math.Sqrt(math.Pow(currentLocation.Latitude-destination.Latitude, 2) + math.Pow(currentLocation.Longitude-destination.Longitude, 2))
	return travelDuration
}
