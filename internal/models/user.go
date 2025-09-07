package models

import (
	"errors"
	"log"
	"strings"

	"github.com/alfandevt/go-event-api/internal/database"
	"github.com/alfandevt/go-event-api/pkg/utils"
)

type User struct {
	ID       int64
	Email    string
	Username string
	Password string
}

// DTOs
type SignUpDTO struct {
	Email    string `binding:"required"`
	Username string `binding:"required"`
	Password string `binding:"required"`
}
type SignInDTO struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	q := `INSERT INTO users (email, username, password) VALUES (?,?,?);`
	stmt, err := database.DB.Prepare(q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPass, err := utils.Hash(u.Password)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		strings.ToLower(u.Email),
		strings.ToLower(u.Username),
		hashedPass)

	return err
}

func (u *User) ValidateCredentials() error {
	q := `SELECT id, password FROM users WHERE email = ?;`
	row := database.DB.QueryRow(q, u.Email)

	var hashedPass string
	err := row.Scan(&u.ID, &hashedPass)
	if err != nil {
		log.Println(err)
		return errors.New("invalid credentials")
	}

	passwordMatch := utils.CheckPaswordHash(u.Password, hashedPass)
	if !passwordMatch {
		return errors.New("invalid credentials")
	}

	return nil
}
