package controllers

import (
	"api-order/src/alert/application"                 // Adjusted import path
	"api-order/src/alert/infrastructure/http/request" // Adjusted import path
	"api-order/src/shared/responses"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RegisterAlertController struct {
	AlertService *application.RegisterAlertUseCase
	Validator    *validator.Validate
}

func NewRegisterAlertController(service *application.RegisterAlertUseCase) *RegisterAlertController {
	return &RegisterAlertController{
		AlertService: service,
		Validator:    validator.New(),
	}
}

// @Summary      Register a new alert
// @Description  Creates a new alert record for a specific kit.
// @Tags         Alerts
// @Accept       json
// @Produce      json
// @Param        alert body request.RegisterAlertRequest true "Alert data to register"
// @Security     BearerAuth
// @Success      201  {object}  responses.Response{data=entities.Alert} "Alert registered successfully"
// @Failure      400  {object}  responses.Response "Invalid request body, validation failed, invalid Kit ID, or invalid alert type"
// @Failure      401  {object}  responses.Response "Unauthorized (token missing or invalid)"
// @Failure      500  {object}  responses.Response "Internal server error while registering alert"
// @Router       /v1/alerts/ [post]
func (ctr *RegisterAlertController) Run(ctx *gin.Context) {
	var req request.RegisterAlertRequest

	// 1. Bind JSON request body
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding RegisterAlertRequest: %v", err)
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Invalid request body format.",
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}

	// 2. Validate request struct fields (includes alert_type via 'oneof')
	if err := ctr.Validator.Struct(req); err != nil {
		log.Printf("Validation failed for RegisterAlertRequest: %v", err)
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Validation failed. Check kit_id, alert_type, and message.",
			Data:    nil,
			Error:   err.Error(), // Provides details on which validation failed
		})
		return
	}

	// 3. Call the Use Case
	// Note: Use case already validates alertType internally, but validator catches it earlier.
	createdAlert, err := ctr.AlertService.Run(req.KitID, req.AlertType, req.Message)
	if err != nil {
		log.Printf("Error registering alert for kit %d: %v", req.KitID, err)

		// Check for foreign key constraint error
		if strings.Contains(err.Error(), "FOREIGN KEY constraint failed") ||
			strings.Contains(err.Error(), "1452") { // MySQL FK error
			ctx.JSON(http.StatusBadRequest, responses.Response{
				Success: false,
				Message: "Failed to register alert: Invalid Kit ID.",
				Error:   "The provided kit_id does not exist.",
				Data:    nil,
			})
			return
		}
		// Handle specific use case validation error
		if err.Error() == "invalid alert_type provided" {
			ctx.JSON(http.StatusBadRequest, responses.Response{
				Success: false,
				Message: "Invalid alert type provided.",
				Error:   err.Error(),
				Data:    nil,
			})
			return
		}

		// Generic internal server error
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Failed to register alert.",
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}

	// 4. Return Success Response
	ctx.JSON(http.StatusCreated, responses.Response{
		Success: true,
		Message: "Alert registered successfully.",
		Data:    createdAlert,
		Error:   nil,
	})
}
