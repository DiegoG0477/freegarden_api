package ports

import "api-order/src/data/temperature/domain/entities"

// Interface for temperature data repository operations
type ITemperatureData interface {
	// Creates a new temperature/humidity record
	Create(data entities.TemperatureData) (entities.TemperatureData, error)
	// Potentially add other methods later if needed (e.g., GetByKitID, GetByID)
}
