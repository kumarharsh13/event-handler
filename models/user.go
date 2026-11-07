package models

import (
	"event-handler/db"
	"event-handler/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name" binding:"required"`
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
