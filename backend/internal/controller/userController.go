package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markHiarley/projetinho/internal/model"
	"github.com/markHiarley/projetinho/internal/usecase"
)

type UserController struct {
	LoginController usecase.UserUseCase
}

func NewUserController(usecase usecase.UserUseCase) UserController {
	return UserController{
		LoginController: usecase,
	}
}

func (uc UserController) CreateUser(ctx *gin.Context) {
	var user model.User

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "campos invalidos",
			"details": err.Error(),
		})
		return
	}

	err := uc.LoginController.CreateUser(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "usuario criado com sucesso",
	})

}
