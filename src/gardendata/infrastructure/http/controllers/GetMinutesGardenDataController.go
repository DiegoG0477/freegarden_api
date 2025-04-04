package controllers

import (
	"api-order/src/gardendata/application" // Corrected path
	"api-order/src/gardendata/domain/entities"
	"api-order/src/shared/responses"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetMinutesGardenDataController struct {
	GetUseCase *application.GetMinutesGardenDataUseCase
}

func NewGetMinutesGardenDataController(useCase *application.GetMinutesGardenDataUseCase) *GetMinutesGardenDataController {
	return &GetMinutesGardenDataController{GetUseCase: useCase}
}

// @Summary      Get Recent Garden Data
// @Description  Retrieves garden sensor data records for a specific kit recorded within the last N minutes.
// @Tags         GardenData
// @Produce      json
// @Param        Authorization header string true "Bearer Token or API Key"
// @Param        kit_id   path      int  true  "Kit ID" Format(int64)
// @Param        minutes  path      int  true  "Number of past minutes to fetch data for"
// @Success      200      {object}  responses.Response{data=[]entities.GardenDataResponse} "Data retrieved successfully"
// @Failure      400      {object}  responses.Response "Invalid Kit ID or Minutes parameter"
// @Failure      401      {object}  responses.Response "Unauthorized - Invalid or missing token/key"
// @Failure      403      {object}  responses.Response "Forbidden - User does not have access to this Kit ID"
// @Failure      404      {object}  responses.Response "Kit ID not found (if validation added)"
// @Failure      500      {object}  responses.Response "Internal server error while retrieving data"
// @Router       /v1/garden/data/kit/{kit_id}/minutes/{minutes} [get]
// @Security     BearerAuth // Or appropriate scheme
func (ctr *GetMinutesGardenDataController) Run(ctx *gin.Context) {
	// Parse Kit ID from path
	kitIDParam := ctx.Param("kit_id")
	kitID, err := strconv.ParseInt(kitIDParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false, Message: "ID de kit inválido.", Error: err.Error(), Data: nil,
		})
		return
	}

	// Parse Minutes from path
	minutesParam := ctx.Param("minutes")
	minutes, err := strconv.Atoi(minutesParam)
	if err != nil || minutes <= 0 { // Also check if minutes is positive
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false, Message: "Parámetro de minutos inválido (debe ser un entero positivo).", Error: err.Error(), Data: nil,
		})
		return
	}

	// Optional: Authorization Check
	// Ensure the authenticated user (from ctx, set by middleware) has access to this kitID.
	// userIDFromToken, exists := ctx.Get("userID")
	// if !exists { ... handle unauthorized ... }
	// hasAccess := checkUserKitAccess(userIDFromToken, kitID) // Implement this logic
	// if !hasAccess { ctx.JSON(http.StatusForbidden, ...); return }

	// Execute the use case
	records, err := ctr.GetUseCase.Run(kitID, minutes)

	// Handle errors from use case
	if err != nil {
		// Log the internal error
		// log.Printf("Error getting garden data for kit %d (%d min): %v", kitID, minutes, err)

		// Check for specific errors like "not found" if the repository/use case signals it
		if errors.Is(err, sql.ErrNoRows) { // Or a custom "NotFound" error
			ctx.JSON(http.StatusNotFound, responses.Response{
				Success: false, Message: "No se encontraron datos para el kit especificado o el kit no existe.", Error: "Not found", Data: nil,
			})
		} else {
			// Generic internal server error
			ctx.JSON(http.StatusInternalServerError, responses.Response{
				Success: false, Message: "Error al obtener los datos del jardín.", Error: "Internal server error", Data: nil,
			})
		}
		return
	}

	// Convert records to response format (if different, though here it's the same)
	responseRecords := make([]entities.GardenDataResponse, len(records))
	for i, rec := range records {
		responseRecords[i] = rec.ToResponse()
	}

	// Return success response (even if the list is empty)
	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: fmt.Sprintf("Datos de los últimos %d minutos obtenidos correctamente.", minutes),
		Data:    responseRecords, // Send the slice of response objects
		Error:   nil,
	})
}
