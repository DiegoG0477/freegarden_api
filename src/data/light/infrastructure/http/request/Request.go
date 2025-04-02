package request

// Request struct for registering a light level record
type RegisterLightRequest struct {
	KitID      int     `json:"kit_id" validate:"required,gt=0"`       // Kit ID must be provided and positive
	LightLevel float64 `json:"light_level" validate:"required,gte=0"` // Light level must be provided and non-negative (adjust if needed)
}
