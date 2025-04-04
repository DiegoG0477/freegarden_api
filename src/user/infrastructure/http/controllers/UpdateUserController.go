package controllers

import (
	"api-order/src/shared/responses"
	"api-order/src/user/application"
	"api-order/src/user/infrastructure/http/request"
	"database/sql" // Required for sql.ErrNoRows if checking that specifically
	"errors"       // Required for errors.Is
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UpdateUserController struct {
	UserService *application.UpdateUserUseCase
	Validator   *validator.Validate
}

func NewUpdateUserController(userService *application.UpdateUserUseCase) *UpdateUserController {
	return &UpdateUserController{
		UserService: userService,
		Validator:   validator.New(),
	}
}

// @Summary      Update user information
// @Description  Updates the first name and last name for the specified user ID. Requires authentication, and users can typically only update their own data.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer Token"
// @Param        id path int true "User ID" Format(int64)
// @Param        user body request.UpdateUserRequest true "User Data to Update"
// @Success      200  {object}  responses.Response{data=entities.UserResponse} "User updated successfully"
// @Failure      400  {object}  responses.Response "Invalid user ID or request body validation failed"
// @Failure      401  {object}  responses.Response "Unauthorized - Invalid or missing token"
// @Failure      403  {object}  responses.Response "Forbidden - User attempting to update another user's data"
// @Failure      404  {object}  responses.Response "User not found"
// @Failure      500  {object}  responses.Response "Internal server error during update"
// @Router       /v1/users/{id} [put]
// @Security     BearerAuth
func (ctr *UpdateUserController) Run(ctx *gin.Context) {
	// 1. Get User ID from Path Parameter
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false, Message: "ID de usuario inválido.", Error: err.Error(), Data: nil,
		})
		return
	}

	// 2. Authorization Check: Ensure the authenticated user matches the ID being updated
	userIDFromToken, exists := ctx.Get("userID") // Assuming middleware sets "userID" as int64
	if !exists {
		ctx.JSON(http.StatusUnauthorized, responses.Response{
			Success: false, Message: "No autorizado: Falta información de usuario.", Error: "Missing user context", Data: nil,
		})
		return
	}
	// Type assertion to ensure it's int64
	authenticatedUserID, ok := userIDFromToken.(int64)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false, Message: "Error interno del servidor.", Error: "Invalid user ID type in context", Data: nil,
		})
		return
	}

	if authenticatedUserID != id {
		// Optional: Allow admins based on role check here
		// roleFromToken, _ := ctx.Get("role")
		// if roleFromToken != "admin" { ... }
		ctx.JSON(http.StatusForbidden, responses.Response{
			Success: false, Message: "Acceso denegado: No puedes modificar datos de otro usuario.", Error: "Forbidden", Data: nil,
		})
		return
	}

	// 3. Bind and Validate Request Body
	var req request.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false, Message: "Error procesando la solicitud. Verifique los campos.", Error: err.Error(), Data: nil,
		})
		return
	}
	if err := ctr.Validator.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false, Message: "Datos inválidos proporcionados.", Error: err.Error(), Data: nil,
		})
		return
	}

	// 4. Execute Update Use Case
	updatedUser, err := ctr.UserService.Run(id, req.FirstName, req.LastName)

	// 5. Handle Errors
	if err != nil {
		// Check if the error indicates the user wasn't found for update
		if errors.Is(err, sql.ErrNoRows) || strings.Contains(err.Error(), "not found") {
			ctx.JSON(http.StatusNotFound, responses.Response{
				Success: false, Message: "Usuario no encontrado para actualizar.", Error: "User not found", Data: nil,
			})
		} else {
			// Handle other potential errors (DB connection issues, etc.)
			ctx.JSON(http.StatusInternalServerError, responses.Response{
				Success: false, Message: "Error al actualizar el usuario.", Error: "Internal server error", Data: nil,
			})
		}
		return
	}

	// 6. Return Success Response
	ctx.JSON(http.StatusOK, responses.Response{
		Success: true, Message: "Usuario actualizado correctamente.", Error: nil, Data: updatedUser.ToResponse(), // Use response struct
	})
}
