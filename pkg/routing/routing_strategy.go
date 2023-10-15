package routing

import "com.vorto.vehiclerouter/pkg/models"

const (
	MaxTime = 12.0 * 60.0
)

// The RoutingStrategy interface provides an abstraction for different routing algorithms
// aimed at optimizing vehicle routing tasks. Implementations of this interface encapsulate
// specific strategies to determine the most efficient path for drivers to deliver loads.
type RoutingStrategy interface {
	ScheduleLoads(loads []models.Load) []models.DriverDeliveryAssignment
}
