package request

// Request struct for registering an air quality record
type RegisterAirQualityRequest struct {
	KitID           int `json:"kit_id" validate:"required,gt=0"`             // Kit ID must be provided and positive
	AirQualityIndex int `json:"air_quality_index" validate:"required,gte=0"` // AQI must be provided and non-negative (adjust range if needed)
}
