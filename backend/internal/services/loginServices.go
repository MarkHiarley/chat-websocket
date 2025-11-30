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

func (au *LoginService) AuthenticateUser(email string) (senha, role string, error error) {
	query := "SELECT password, role from users WHERE email = $1"

	row := au.connection.QueryRow(query, email)

	err := row.Scan(&senha, &role)
	if err != nil {

		return "", "", err
	}

	return senha, role, nil
}
