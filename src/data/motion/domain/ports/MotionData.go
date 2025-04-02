package ports

import "api-order/src/data/motion/domain/entities" // Adjusted import path

// Interface for motion data repository operations
type IMotionData interface {
	// Creates a new motion detection record
	Create(data entities.MotionData) (entities.MotionData, error)
	// Add other methods later if needed
}
