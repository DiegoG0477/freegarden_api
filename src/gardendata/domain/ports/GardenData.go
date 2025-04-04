package ports

import "api-order/src/gardendata/domain/entities" // Corrected path

// IGardenData defines the interface for the garden data repository.
type IGardenData interface {
	// Create saves a new garden data record to the repository.
	Create(data entities.GardenData) (entities.GardenData, error)

	// GetRecordsByKitIDAndTime retrieves records for a specific kit within a given time window (in minutes).
	GetRecordsByKitIDAndTime(kitID int64, minutesAgo int) ([]entities.GardenData, error)
}
