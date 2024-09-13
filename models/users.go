package models

import (
	"errors"

	"github.com/Soyaib10/eba-event-booking-api/db"
	"github.com/Soyaib10/eba-event-booking-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrivedPassword string
	err := row.Scan(&u.ID, &retrivedPassword)
	if err != nil {
		return errors.New("credentials invalid")
	}
	isPasswordValid := utils.CheckPasswordHash(u.Password, retrivedPassword)
	if !isPasswordValid {
		return errors.New("credentials invalid")
	}
	return nil
}

func (u User) Save() error {
	// Insert statement and prepare it
	qurey := `INSERT INTO users(email, password)
			VALUES(?, ?)`
	stmt, err := db.DB.Prepare(qurey)
	if err != nil {
		return err
	}
	defer stmt.Close()
	
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	// Execute stmt
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}