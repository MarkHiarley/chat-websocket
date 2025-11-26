package services

import (
	"database/sql"

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
	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)"

	userName := user.Username
	userEmail := user.Email

	_, err := us.connection.Exec(query, userName, userEmail, senhaHash)

	if err != nil {
		return err
	}

	return nil
}
