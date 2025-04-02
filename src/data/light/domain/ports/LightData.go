package ports

import "api-order/src/data/light/domain/entities" // Adjusted import path

// Interface for light data repository operations
type ILightData interface {
	// Creates a new light level record
	Create(data entities.LightData) (entities.LightData, error)
	// Add other methods later if needed
}
