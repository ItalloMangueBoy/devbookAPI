package helper

import (
	db "devbookAPI/src/database"
	"devbookAPI/src/model"
	"errors"
)

// CreateUser: insert one user in database
func CreateUser(user *model.User) error {
	stmt, err := db.Conn.Prepare("INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return errors.New("cannot generate insert querry")
	}

	res, err := stmt.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return err
	}

	user.Id, err = res.LastInsertId()
	if err != nil {
		return errors.New("cannot search inserted id into database")
	}

	return nil
}


