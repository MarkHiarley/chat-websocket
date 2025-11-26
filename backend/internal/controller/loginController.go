package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markHiarley/projetinho/internal/model"
	"github.com/markHiarley/projetinho/internal/usecase"
)

type LoginController struct {
	LoginUseCase usecase.LoginUseCase
}

func NewLoginController(usecase usecase.LoginUseCase) LoginController {
	return LoginController{
		LoginUseCase: usecase,
	}
}

func (au LoginController) AuthenticateUser(ctx *gin.Context) {
	var login model.LoginUser

	if err := ctx.ShouldBind(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "campos invalidos",
			"details": err.Error(),
		})
		return
	}

	accessToken, refreshToken, err := au.LoginUseCase.AuthenticateUser(login)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})

}
