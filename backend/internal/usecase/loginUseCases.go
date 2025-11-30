package usecase

import (
	"github.com/markHiarley/projetinho/internal/auth"
	"github.com/markHiarley/projetinho/internal/model"
	"github.com/markHiarley/projetinho/internal/services"
	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase struct {
	services services.LoginService
}

func NewLoginUseCase(services services.LoginService) LoginUseCase {
	return LoginUseCase{
		services: services,
	}
}

func (au *LoginUseCase) AuthenticateUser(body model.LoginUser) (accessToken,
	refreshToken string, error error) {
	email := body.Email.String

	senhaLogin := body.Password.String

	senhaHash, role, err := au.services.AuthenticateUser(email)

	if err != nil {
		return "", "", err
	}

	comparado := bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senhaLogin))

	if comparado != nil {
		return "", "", comparado
	}

	accessToken1, err := auth.GenerateAccessToken(email, role)
	if err != nil {
		return "", "", err
	}
	refreshToken1, err := auth.GenerateRefreshToken(email, role)
	if err != nil {
		return "", "", err
	}

	return accessToken1, refreshToken1, nil

}
