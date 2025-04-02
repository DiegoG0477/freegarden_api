package controllers

import (
	"api-order/src/data/motion/application"                 // Adjusted import path
	"api-order/src/data/motion/infrastructure/http/request" // Adjusted import path
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
	var req request.RegisterMotionRequest

	// 1. Bind JSON request body
	// ShouldBindJSON will require motion_detected to be present (true or false)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding RegisterMotionRequest: %v", err)
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Invalid request body format. Ensure 'kit_id' (number) and 'motion_detected' (boolean) are provided.",
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}

	// 2. Validate request struct fields (primarily kit_id's gt=0)
	if err := ctr.Validator.Struct(req); err != nil {
		log.Printf("Validation failed for RegisterMotionRequest: %v", err)
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Validation failed. Check kit_id.",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	// 3. Call the Use Case with data from request
	createdRecord, err := ctr.RecordService.Run(req.KitID, req.MotionDetected)
	if err != nil {
		log.Printf("Error registering motion record for kit %d: %v", req.KitID, err)

		// Specific check for foreign key constraint violation
		if strings.Contains(err.Error(), "FOREIGN KEY constraint failed") ||
			strings.Contains(err.Error(), "1452") { // MySQL foreign key error code
			ctx.JSON(http.StatusBadRequest, responses.Response{
				Success: false,
				Message: "Failed to register record: Invalid Kit ID.",
				Error:   "The provided kit_id does not exist.",
				Data:    nil,
			})
			return
		}

		// Generic internal server error
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Failed to register motion record.",
			Error:   err.Error(), // Avoid exposing raw DB errors in production
			Data:    nil,
		})
		return
	}

	// 4. Return Success Response
	ctx.JSON(http.StatusCreated, responses.Response{
		Success: true,
		Message: "Motion record registered successfully.",
		Data:    createdRecord,
		Error:   nil,
	})
}
