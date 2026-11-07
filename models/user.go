package models

import (
	"errors"
	"event-handler/db"
	"event-handler/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"-" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	password, err := utils.HashPassword(u.Password)
	u.Password = password
	if err != nil {
		return err
	}

	query := `
        INSERT INTO users (name, email, password)
        VALUES($1, $2, $3)
        RETURNING id
    `
	err = db.DB.QueryRow(query, u.Name, u.Email, u.Password).Scan(&u.ID)

	return err
}

func (login *LoginRequest) ValidateCredentials() error {
	var retrievedPassword string
	query := `SELECT password FROM users WHERE email = $1`
	err := db.DB.QueryRow(query, login.Email).Scan(&retrievedPassword)

	if err != nil {
		return errors.New("Credentials are invalid")
	}

	passwordValid := utils.CheckPasswordHash(login.Password, retrievedPassword)

	if !passwordValid {
		return errors.New("Credentials are invalid")
	}

	return nil
}
