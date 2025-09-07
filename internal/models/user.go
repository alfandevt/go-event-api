package models

import "github.com/alfandevt/go-even-api/internal/database"

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Username string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	q := `INSERT INTO users (email, username, password) VALUES (?,?,?);`
	stmt, err := database.DB.Prepare(q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(u.Email, u.Username, u.Password)
	if err != nil {
		return err
	}

	return nil
}
