package controllers

import (
	"api-order/src/kit/application"
	"api-order/src/kit/infrastructure/http/request"
	"api-order/src/shared/middlewares" // Import middleware package
	"api-order/src/shared/responses"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateKitController struct {
	KitService *application.CreateKitUseCase
	Validator  *validator.Validate
}

func NewCreateKitController(kitService *application.CreateKitUseCase) *CreateKitController {
	return &CreateKitController{
		KitService: kitService,
		Validator:  validator.New(),
	}
}

func (ctr *CreateKitController) Run(ctx *gin.Context) {
	var req request.CreateKitRequest

	// 1. Bind JSON request body
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding CreateKitRequest: %v", err)
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
		log.Printf("Validation failed for CreateKitRequest: %v", err)
		// Provide more specific validation errors if needed
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Validation failed.",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	// 3. Get UserID from JWT Claims (set by middleware)
	claimsData, exists := ctx.Get("datUser")
	if !exists {
		log.Println("Error: datUser claims not found in context. Middleware might not have run.")
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
		log.Println("Error: Failed to assert datUser claims to *middlewares.CustomClaims.")
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Internal Server Error: Could not process user identity.",
			Error:   "Type assertion failed for claims.",
			Data:    nil,
		})
		return
	}
	userID := customClaims.ClientID // Get the user ID

	// 4. Call the Use Case
	createdKit, err := ctr.KitService.Run(req.Name, req.Description, userID)
	if err != nil {
		log.Printf("Error creating kit for user %d: %v", userID, err)
		// Handle specific errors, e.g., foreign key constraints if user_id is invalid, etc.
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Failed to create kit.",
			Error:   err.Error(), // Be cautious about exposing raw DB errors
			Data:    nil,
		})
		return
	}

	// 5. Return Success Response
	ctx.JSON(http.StatusCreated, responses.Response{
		Success: true,
		Message: "Kit created successfully.",
		Data:    createdKit,
		Error:   nil,
	})
}
