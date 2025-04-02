package ports

import "api-order/src/data/airquality/domain/entities" // Adjusted import path

// Interface for air quality data repository operations
type IAirQualityData interface {
	// Creates a new air quality record
	Create(data entities.AirQualityData) (entities.AirQualityData, error)
	// Add other methods later if needed
}
