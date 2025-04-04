package request

// RegisterGardenDataRequest defines the expected JSON body for registering new data.
type RegisterGardenDataRequest struct {
	KitID               int64   `json:"kit_id" validate:"required,gt=0"` // Ensure KitID is positive
	Temperature         float64 `json:"temperature"`                     // Add validation tags if needed (e.g., min/max)
	GroundHumidity      float64 `json:"ground_humidity"`
	EnvironmentHumidity float64 `json:"environment_humidity"` // Corrected spelling
	PhLevel             float64 `json:"ph_level"`
	Time                int64   `json:"time" validate:"required"` // Require the device timestamp
}

// Note: No specific request struct is typically needed for the GET request,
// as parameters are usually passed via URL path or query strings.
