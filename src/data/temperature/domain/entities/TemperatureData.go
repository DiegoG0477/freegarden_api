package entities

import "time"

// Represents a single temperature and humidity reading
type TemperatureData struct {
	DataID      int       `json:"data_id"`     // Corresponds to data_id PK
	KitID       int       `json:"kit_id"`      // Foreign key to kits table
	Temperature float64   `json:"temperature"` // Temperature reading
	Humidity    float64   `json:"humidity"`    // Humidity reading
	Timestamp   time.Time `json:"timestamp"`   // Timestamp from DB default
}
