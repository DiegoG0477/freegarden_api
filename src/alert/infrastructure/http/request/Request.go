package request

// Request struct for registering an alert
// Use `oneof` validator tag for alert_type based on defined constants
type RegisterAlertRequest struct {
	KitID     int    `json:"kit_id" validate:"required,gt=0"`
	AlertType string `json:"alert_type" validate:"required,oneof=under_min higher_max"`
	Message   string `json:"message" validate:"required,min=1"`
}

// No request body for GetAlertsByKitID
