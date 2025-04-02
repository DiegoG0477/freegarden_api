package controllers

import (
	"api-order/src/data/temperature/application"
	"api-order/src/data/temperature/infrastructure/http/request"
	"api-order/src/shared/responses"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RegisterRecordController struct {
	RecordService *application.RegisterRecordUseCase
	Validator     *validator.Validate
}

func NewRegisterRecordController(service *application.RegisterRecordUseCase) *RegisterRecordController {
	return &RegisterRecordController{
		RecordService: service,
		Validator:     validator.New(),
	}
}

func (ctr *RegisterRecordController) Run(ctx *gin.Context) {
	var req request.RegisterRecordRequest

	// 1. Bind JSON request body
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding RegisterRecordRequest: %v", err)
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Invalid request body format.",
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}

	// 2. Validate request struct fields
	if err := ctr.Validator.Struct(req); err != nil {
		log.Printf("Validation failed for RegisterRecordRequest: %v", err)
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Validation failed. Check kit_id, temperature, and humidity.",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	// 3. Call the Use Case with data from request
	// Note: We are NOT getting kit_id from JWT here, as per requirements.
	createdRecord, err := ctr.RecordService.Run(req.KitID, req.Temperature, req.Humidity)
	if err != nil {
		log.Printf("Error registering temperature record for kit %d: %v", req.KitID, err)

		// Specific check for foreign key constraint violation (example)
		if strings.Contains(err.Error(), "FOREIGN KEY constraint failed") || // SQLite / some DBs
			strings.Contains(err.Error(), "1452") { // MySQL foreign key error code
			ctx.JSON(http.StatusBadRequest, responses.Response{
				Success: false,
				Message: "Failed to register record: Invalid Kit ID.",
				Error:   "The provided kit_id does not exist.",
				Data:    nil,
			})
			return
		}

		// Generic internal server error for other issues
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Failed to register temperature record.",
			Error:   err.Error(), // Be cautious exposing raw DB errors in production
			Data:    nil,
		})
		return
	}

	// 4. Return Success Response
	ctx.JSON(http.StatusCreated, responses.Response{
		Success: true,
		Message: "Temperature record registered successfully.",
		Data:    createdRecord, // Contains DataID, KitID, Temp, Humidity (Timestamp is zero/nil here)
		Error:   nil,
	})
}
