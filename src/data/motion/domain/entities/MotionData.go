package entities

import "time"

// Represents a single motion detection event
type MotionData struct {
	DataID         int       `json:"data_id"`         // Corresponds to data_id PK
	KitID          int       `json:"kit_id"`          // Foreign key to kits table
	MotionDetected bool      `json:"motion_detected"` // Whether motion was detected
	Timestamp      time.Time `json:"timestamp"`       // Timestamp from DB default
}
