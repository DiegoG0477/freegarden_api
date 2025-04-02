package ports

import "api-order/src/alert/domain/entities" // Adjusted import path

// Interface for alert repository operations
type IAlert interface {
	// Creates a new alert record
	Create(alert entities.Alert) (entities.Alert, error)
	// Retrieves all alerts associated with a specific kit ID
	GetByKitID(kitID int) ([]entities.Alert, error)
}
