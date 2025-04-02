package request

// Request struct for registering a temperature/humidity record
type RegisterRecordRequest struct {
	KitID       int     `json:"kit_id" validate:"required,gt=0"` // Kit ID must be provided and positive
	Temperature float64 `json:"temperature" validate:"required"` // Use appropriate validation (e.g., numeric range) if needed
	Humidity    float64 `json:"humidity" validate:"required"`    // Use appropriate validation (e.g., numeric range 0-100) if needed
}
