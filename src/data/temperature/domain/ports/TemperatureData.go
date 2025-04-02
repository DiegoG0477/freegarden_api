package ports

import "api-order/src/data/temperature/domain/entities"

// Interface for temperature data repository operations
type ITemperatureData interface {
	// Creates a new temperature/humidity record
	Create(data entities.TemperatureData) (entities.TemperatureData, error)

	// Retrieves records for a specific kit within the last N minutes
	GetRecordsByKitIDAndTime(kitID int, minutesAgo int) ([]entities.TemperatureData, error) // <- NUEVO METODO
}
