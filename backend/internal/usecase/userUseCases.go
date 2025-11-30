package usecase

import (
	"fmt"

	"github.com/markHiarley/projetinho/internal/model"
	"github.com/markHiarley/projetinho/internal/services"
	"golang.org/x/crypto/bcrypt"
)

var JWT_SECRET_TOKEN []byte

type UserUseCase struct {
	UserService services.UserService
}

func NewUserUseCase(service services.UserService) UserUseCase {
	return UserUseCase{
		UserService: service,
	}
}

func (uc UserUseCase) CreateUser(user model.User) error {
	senhaHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	senhaHashToString := string(senhaHash)

	err = uc.UserService.CreateUser(user, senhaHashToString)

	if err != nil {
		return fmt.Errorf("email ja cadastrado, tente usar outro email: %w", err)
	}

	return nil
}
func (uc UserUseCase) DeleteUser(id string) error {

	err := uc.UserService.DeleteUser(id)

	if err != nil {
		return fmt.Errorf("erro ao deletar usuario: %w", err)
	}

	return nil
}
