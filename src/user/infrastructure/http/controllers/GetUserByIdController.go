package controllers

import (
	"api-order/src/shared/responses"
	"api-order/src/user/application"
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type GetUserByIdController struct {
	UserService *application.GetUserByIdUseCase
}

func NewGetUserByIdController(userService *application.GetUserByIdUseCase) *GetUserByIdController {
	return &GetUserByIdController{UserService: userService}
}

// @Summary      Get user by ID
// @Description  Retrieves the details of a specific user by their ID. Requires authentication.
// @Tags         Users
// @Produce      json
// @Param        Authorization header string true "Bearer Token"
// @Param        id path int true "User ID" Format(int64)
// @Success      200  {object}  responses.Response{data=entities.UserResponse} "User retrieved successfully"
// @Failure      400  {object}  responses.Response "Invalid user ID provided"
// @Failure      401  {object}  responses.Response "Unauthorized - Invalid or missing token"
// @Failure      403  {object}  responses.Response "Forbidden - User attempting to access another user's data (if implemented)"
// @Failure      404  {object}  responses.Response "User not found"
// @Failure      500  {object}  responses.Response "Internal server error while retrieving user"
// @Router       /v1/users/{id} [get]
// @Security     BearerAuth
func (ctr *GetUserByIdController) Run(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false, Message: "ID de usuario inválido.", Error: err.Error(), Data: nil,
		})
		return
	}

	// Optional: Authorization Check (ensure user can only get their own data unless admin)
	// userIDFromToken, exists := ctx.Get("userID") // Assuming middleware sets this
	// if !exists || userIDFromToken.(int64) != id {
	//     // You might allow admins to bypass this check based on a role claim in the token
	//     ctx.JSON(http.StatusForbidden, responses.Response{
	//         Success: false, Message: "Acceso denegado.", Error: "Forbidden", Data: nil,
	//     })
	//     return
	// }

	user, err := ctr.UserService.Run(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || strings.Contains(err.Error(), "not found") {
			ctx.JSON(http.StatusNotFound, responses.Response{
				Success: false, Message: "Usuario no encontrado.", Error: "User not found", Data: nil,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, responses.Response{
				Success: false, Message: "Error al obtener el usuario.", Error: "Internal server error", Data: nil,
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true, Message: "Usuario obtenido con éxito.", Error: nil, Data: user.ToResponse(), // Return response without password
	})
}
