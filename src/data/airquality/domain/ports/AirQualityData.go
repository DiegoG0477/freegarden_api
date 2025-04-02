package ports

import "api-order/src/data/airquality/domain/entities"

// Interface for air quality data repository operations
type IAirQualityData interface {
	// Creates a new air quality record
	Create(data entities.AirQualityData) (entities.AirQualityData, error)

	// Retrieves air quality records for a specific kit within the last N minutes
	GetRecordsByKitIDAndTime(kitID int, minutesAgo int) ([]entities.AirQualityData, error) // <- NUEVO METODO
}
