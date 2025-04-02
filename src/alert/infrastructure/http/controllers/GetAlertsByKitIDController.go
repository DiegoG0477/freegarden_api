package controllers

import (
	"api-order/src/alert/application" // Adjusted import path
	"api-order/src/shared/responses"
	"log"
	"net/http"
	"strconv" // For parsing kit_id from URL

	"github.com/gin-gonic/gin"
)

type GetAlertsByKitIDController struct {
	AlertService *application.GetAlertsByKitIDUseCase
}

func NewGetAlertsByKitIDController(service *application.GetAlertsByKitIDUseCase) *GetAlertsByKitIDController {
	return &GetAlertsByKitIDController{AlertService: service}
}

func (ctr *GetAlertsByKitIDController) Run(ctx *gin.Context) {
	// 1. Get kit_id from URL parameter
	kitIDParam := ctx.Param("kit_id")
	kitID, err := strconv.Atoi(kitIDParam) // Use Atoi for int conversion

	if err != nil || kitID <= 0 { // Also check if kitID is positive
		log.Printf("Invalid kit_id parameter received: %s", kitIDParam)
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Invalid Kit ID provided in URL.",
			Data:    nil,
			Error:   "Kit ID must be a positive integer.",
		})
		return
	}

	// 2. Call the Use Case
	alerts, err := ctr.AlertService.Run(kitID)
	if err != nil {
		// Log error, but don't necessarily expose DB details
		log.Printf("Error retrieving alerts for kit %d: %v", kitID, err)
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Failed to retrieve alerts.",
			Error:   "An internal error occurred.", // Keep error message generic
			Data:    nil,
		})
		return
	}

	// 3. Return Success Response (even if alerts slice is empty)
	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "Alerts retrieved successfully.",
		Data:    alerts, // Will be [] if no alerts found
		Error:   nil,
	})
}
