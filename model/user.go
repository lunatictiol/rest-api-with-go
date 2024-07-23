package model

import (
	"errors"

	"github.comlunatictiol/rest-api-with-go/db"
	"github.comlunatictiol/rest-api-with-go/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {

	query := "INSERT INTO users(email,password) VALUES(?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		println("1", err)
		return err
	}
	defer stmt.Close()
	pswrd, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, pswrd)
	if err != nil {
		println("3", err)
		return err
	}
	userId, err := result.LastInsertId()

	if err != nil {
		println("2", err)
		return err
	}
	u.Id = userId
	return nil
}

func (u *User) Validate() error {

	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.Id, &retrievedPassword)
	if err != nil {
		return errors.New("invalid password or email")
	}
	passwordIsValid := utils.ValidatePassword(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("invalid password")
	}
	return nil

}
