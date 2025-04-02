package entities

import "time"

// Define constants for known alert types for validation and consistency
const (
	AlertTypeUnderMin  = "under_min"
	AlertTypeHigherMax = "higher_max"
	// Add other alert types here if needed
)

// IsValidAlertType checks if a given string is a valid alert type
func IsValidAlertType(alertType string) bool {
	switch alertType {
	case AlertTypeUnderMin, AlertTypeHigherMax:
		return true
	default:
		return false
	}
}

// Represents a single alert record
type Alert struct {
	AlertID   int       `json:"alert_id"`   // Corresponds to alert_id PK
	KitID     int       `json:"kit_id"`     // Foreign key to kits table
	AlertType string    `json:"alert_type"` // Type of alert (e.g., "under_min", "higher_max")
	Message   string    `json:"message"`    // Detailed message for the alert
	Timestamp time.Time `json:"timestamp"`  // Timestamp from DB default
}
