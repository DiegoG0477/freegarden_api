package entities

import "time"

// Represents a single light level reading
type LightData struct {
	DataID     int       `json:"data_id"`     // Corresponds to data_id PK
	KitID      int       `json:"kit_id"`      // Foreign key to kits table
	LightLevel float64   `json:"light_level"` // Light level reading (e.g., lux)
	Timestamp  time.Time `json:"timestamp"`   // Timestamp from DB default
}
