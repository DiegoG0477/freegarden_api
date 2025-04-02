package ports

import "api-order/src/data/motion/domain/entities"

// Interface for motion data repository operations
type IMotionData interface {
	// Creates a new motion detection record
	Create(data entities.MotionData) (entities.MotionData, error)

	// Retrieves motion records for a specific kit within the last N minutes
	GetRecordsByKitIDAndTime(kitID int, minutesAgo int) ([]entities.MotionData, error) // <- NUEVO METODO
}
