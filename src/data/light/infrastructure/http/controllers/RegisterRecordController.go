package controllers

import (
	"api-order/src/data/light/application"                 // Adjusted import path
	"api-order/src/data/light/infrastructure/http/request" // Adjusted import path
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

// @Summary      Register light data record
// @Description  Creates a new light level data record for a specific kit.
// @Tags         Light Data
// @Accept       json
// @Produce      json
// @Param        lightData body request.RegisterLightRequest true "Light data record"
// @Security     BearerAuth
// @Success      201  {object}  responses.Response{data=entities.LightData} "Light record registered successfully"
// @Failure      400  {object}  responses.Response "Invalid request body, validation failed, or invalid Kit ID"
// @Failure      401  {object}  responses.Response "Unauthorized (token missing or invalid)"
// @Failure      500  {object}  responses.Response "Internal server error while registering light record"
// @Router       /data/light/ [post]
func (ctr *RegisterRecordController) Run(ctx *gin.Context) {
	var req request.RegisterLightRequest

	// 1. Bind JSON request body
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding RegisterLightRequest: %v", err)
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
		log.Printf("Validation failed for RegisterLightRequest: %v", err)
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Validation failed. Check kit_id and light_level.",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	// 3. Call the Use Case with data from request
	createdRecord, err := ctr.RecordService.Run(req.KitID, req.LightLevel)
	if err != nil {
		log.Printf("Error registering light record for kit %d: %v", req.KitID, err)

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
			Message: "Failed to register light record.",
			Error:   err.Error(), // Avoid exposing raw DB errors in production
			Data:    nil,
		})
		return
	}

	// 4. Return Success Response
	ctx.JSON(http.StatusCreated, responses.Response{
		Success: true,
		Message: "Light record registered successfully.",
		Data:    createdRecord,
		Error:   nil,
	})
}
