package entities

import "time"

// Represents a single air quality reading
type AirQualityData struct {
	DataID          int       `json:"data_id"`           // Corresponds to data_id PK
	KitID           int       `json:"kit_id"`            // Foreign key to kits table
	AirQualityIndex int       `json:"air_quality_index"` // Air Quality Index reading
	Timestamp       time.Time `json:"timestamp"`         // Timestamp from DB default
}
