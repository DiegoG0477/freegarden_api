package controllers

import (
	"api-order/src/data/light/application" // Ajustado
	"api-order/src/shared/responses"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetMinutesRecordsController struct {
	RecordService *application.GetMinutesRecordsUseCase // Ajustado
}

func NewGetMinutesRecordsController(service *application.GetMinutesRecordsUseCase) *GetMinutesRecordsController { // Ajustado
	return &GetMinutesRecordsController{
		RecordService: service,
	}
}

// @Summary      Get recent light data records
// @Description  Retrieves light data records for a specific kit within the last N minutes.
// @Tags         Light Data
// @Produce      json
// @Param        kit_id path int true "Kit ID" Format(int64)
// @Param        minutes path int true "Number of minutes ago to fetch records from" Format(int64)
// @Security     BearerAuth
// @Success      200  {object}  responses.Response{data=[]entities.LightData} "Light records retrieved successfully"
// @Failure      400  {object}  responses.Response "Invalid Kit ID or minutes parameter provided"
// @Failure      401  {object}  responses.Response "Unauthorized (token missing or invalid)"
// @Failure      500  {object}  responses.Response "Internal server error while retrieving light records"
// @Router       /data/light/kit/{kit_id}/minutes/{minutes} [get]
func (ctr *GetMinutesRecordsController) Run(ctx *gin.Context) {
	// 1. Get parameters from URL
	kitIDParam := ctx.Param("kit_id")
	minutesParam := ctx.Param("minutes")

	kitID, errKit := strconv.Atoi(kitIDParam)
	minutes, errMinutes := strconv.Atoi(minutesParam)

	// 2. Validate parameters
	if errKit != nil || kitID <= 0 {
		log.Printf("Invalid kit_id parameter received: %s", kitIDParam)
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false, Message: "Invalid Kit ID provided in URL.", Data: nil, Error: "Kit ID must be a positive integer.",
		})
		return
	}
	if errMinutes != nil || minutes <= 0 {
		log.Printf("Invalid minutes parameter received: %s", minutesParam)
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false, Message: "Invalid minutes parameter provided in URL.", Data: nil, Error: "Minutes must be a positive integer.",
		})
		return
	}

	// 3. Call the Use Case
	records, err := ctr.RecordService.Run(kitID, minutes)
	if err != nil {
		log.Printf("Error getting light records for kit %d (%d min): %v", kitID, minutes, err) // Ajustado log
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false, Message: "Failed to retrieve light records.", Error: "An internal error occurred.", Data: nil, // Ajustado mensaje
		})
		return
	}

	// 4. Return Success Response
	ctx.JSON(http.StatusOK, responses.Response{
		Success: true, Message: "Light records retrieved successfully.", Data: records, Error: nil, // Ajustado mensaje
	})
}
