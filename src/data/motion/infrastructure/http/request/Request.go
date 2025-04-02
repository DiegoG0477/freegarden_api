package request

// Request struct for registering a motion detection record
// Using a pointer for bool to distinguish between 'false' and 'not provided' if necessary with validator,
// although ShouldBindJSON typically requires boolean fields to be present.
// Sticking with plain bool for simplicity based on previous examples.
type RegisterMotionRequest struct {
	KitID          int  `json:"kit_id" validate:"required,gt=0"`
	MotionDetected bool `json:"motion_detected"` // No specific validate tag needed, presence checked by ShouldBindJSON
}
