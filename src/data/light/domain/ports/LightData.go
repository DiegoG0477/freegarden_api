package ports

import "api-order/src/data/light/domain/entities"

// Interface for light data repository operations
type ILightData interface {
	// Creates a new light level record
	Create(data entities.LightData) (entities.LightData, error)

	// Retrieves light level records for a specific kit within the last N minutes
	GetRecordsByKitIDAndTime(kitID int, minutesAgo int) ([]entities.LightData, error) // <- NUEVO METODO
}
