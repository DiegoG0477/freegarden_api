package request

// Request struct for creating a kit
type CreateKitRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required"` // Assuming description is mandatory
}

// No request body needed for GetKits by logged-in user
