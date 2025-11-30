package services

import (
	"database/sql"
	"fmt"

	"github.com/markHiarley/projetinho/internal/model"
)

type UserService struct {
	connection *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		connection: db,
	}
}

func (us *UserService) CreateUser(user model.User, senhaHash string) error {
	role := "user"
	query := "INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, $4)"

	userName := user.Username
	userEmail := user.Email

	_, err := us.connection.Exec(query, userName, userEmail, senhaHash, role)

	if err != nil {
		return err
	}

	return nil
}
func (us *UserService) DeleteUser(id string) (error error) {
	query := "DELETE FROM users WHERE id = $1"

	_, err := us.connection.Exec(query, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("usuario não encontrado")
		}
	}

	return nil
}
