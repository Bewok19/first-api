package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	useCase UserService
}

func NewUserController(useCase UserService) *UserController {
	return &UserController{useCase: useCase}
}

func (c UserController) FindAll(ctx *gin.Context) {
	user, err := c.useCase.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}
