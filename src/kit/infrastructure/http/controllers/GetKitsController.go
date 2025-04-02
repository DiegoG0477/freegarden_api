package controllers

import (
	"api-order/src/kit/application"
	"api-order/src/kit/domain/entities"
	"api-order/src/shared/middlewares" // Import middleware package
	"api-order/src/shared/responses"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetKitsController struct {
	KitService *application.GetKitsUseCase
}

func NewGetKitsController(kitService *application.GetKitsUseCase) *GetKitsController {
	return &GetKitsController{KitService: kitService}
}

// @Summary      Get kits for the authenticated user
// @Description  Retrieves all kits associated with the user identified by the JWT token.
// @Tags         Kits
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  responses.Response{data=[]entities.Kit} "Kits retrieved successfully"
// @Failure      401  {object}  responses.Response "Unauthorized"
// @Failure      500  {object}  responses.Response "Internal server error"
// @Router       /kits/ [get]
func (ctr *GetKitsController) Run(ctx *gin.Context) {
	// 1. Get UserID from JWT Claims (set by middleware)
	claimsData, exists := ctx.Get("datUser")
	if !exists {
		log.Println("Error: datUser claims not found in context for GetKits.")
		ctx.JSON(http.StatusUnauthorized, responses.Response{
			Success: false,
			Message: "Unauthorized: User claims not found.",
			Error:   "Authentication context missing.",
			Data:    nil,
		})
		return
	}

	customClaims, ok := claimsData.(*middlewares.CustomClaims)
	if !ok {
		log.Println("Error: Failed to assert datUser claims to *middlewares.CustomClaims for GetKits.")
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Internal Server Error: Could not process user identity.",
			Error:   "Type assertion failed for claims.",
			Data:    nil,
		})
		return
	}
	userID := customClaims.ClientID // Get the user ID

	// 2. Call the Use Case
	kits, err := ctr.KitService.Run(userID)
	if err != nil {
		log.Printf("Error getting kits for user %d: %v", userID, err)
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Failed to retrieve kits.",
			Error:   err.Error(), // Be cautious about exposing raw DB errors
			Data:    nil,
		})
		return
	}

	// 3. Return Success Response
	// Return empty list [] instead of null if no kits found
	if kits == nil {
		kits = []entities.Kit{}
	}
	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "Kits retrieved successfully.",
		Data:    kits,
		Error:   nil,
	})
}
