package services

import (
	"database/sql"
)

type LoginService struct {
	connection *sql.DB
}

func NewLoginService(db *sql.DB) *LoginService {
	return &LoginService{
		connection: db,
	}
}

func (au *LoginService) AuthenticateUser(email string) (senha string, error error) {
	query := "SELECT password from users WHERE email = $1"

	row := au.connection.QueryRow(query, email)

	err := row.Scan(&senha)

	if err != nil {
		return "", err
	}

	return senha, nil
}
