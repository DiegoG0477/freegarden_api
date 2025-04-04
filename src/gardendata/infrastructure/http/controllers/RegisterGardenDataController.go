package controllers

import (
	"api-order/src/gardendata/application"                 // Corrected path
	"api-order/src/gardendata/infrastructure/http/request" // Corrected path
	"api-order/src/shared/responses"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RegisterGardenDataController struct {
	RegisterUseCase *application.RegisterGardenDataUseCase
	Validator       *validator.Validate
}

func NewRegisterGardenDataController(useCase *application.RegisterGardenDataUseCase) *RegisterGardenDataController {
	return &RegisterGardenDataController{
		RegisterUseCase: useCase,
		Validator:       validator.New(),
	}
}

// @Summary      Register Garden Sensor Data
// @Description  Receives and stores a new set of sensor readings for a specific kit.
// @Tags         GardenData
// @Accept       json
// @Produce      json
// @Param        data body request.RegisterGardenDataRequest true "Sensor Data Payload"
// @Success      201  {object}  responses.Response{data=entities.GardenDataResponse} "Data registered successfully"
// @Failure      400  {object}  responses.Response "Invalid request body or validation failed"
// @Failure      401  {object}  responses.Response "Unauthorized - Invalid or missing token/key"
// @Failure      404  {object}  responses.Response "Kit ID not found (if validation added)"
// @Failure      500  {object}  responses.Response "Internal server error during registration"
// @Router       /v1/garden/data/ [post]
func (ctr *RegisterGardenDataController) Run(ctx *gin.Context) {
	var req request.RegisterGardenDataRequest

	// Bind JSON body
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Error procesando la solicitud. Verifique el formato JSON.",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	// Validate struct fields based on tags
	if err := ctr.Validator.Struct(req); err != nil {
		// Log validation errors if needed for debugging
		fmt.Printf("Validation error for RegisterGardenDataRequest: %v\n", err)
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Datos inválidos proporcionados.",
			Data:    nil,
			Error:   err.Error(), // Consider formatting validation errors
		})
		return
	}

	// Optional: Validate KitID existence against the `kits` table or user's associated kits
	// This would likely require injecting a KitRepository or UserService here or in the use case.
	// For now, we assume the kit_id is valid if the DB insert doesn't fail on foreign key.

	// Execute the use case
	createdRecord, err := ctr.RegisterUseCase.Run(
		req.KitID,
		req.Temperature,
		req.GroundHumidity,
		req.EnvironmentHumidity,
		req.PhLevel,
		req.Time,
	)

	if err != nil {
		// Log the error for internal monitoring
		// log.Printf("Error registering garden data via controller: %v", err)

		// Check for specific errors if the use case or repo provides them
		// Example: if errors.Is(err, someSpecificError) { ... }

		// Return a generic server error to the client
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Error al registrar los datos del jardín.",
			Data:    nil,
			Error:   "Internal server error", // Avoid exposing internal error details like err.Error() directly
		})
		return
	}

	// Return success response
	ctx.JSON(http.StatusCreated, responses.Response{
		Success: true,
		Message: "Datos del jardín registrados correctamente.",
		Data:    createdRecord.ToResponse(), // Use the response converter
		Error:   nil,
	})
}
