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
