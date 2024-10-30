package model

import (
	db "devbookAPI/src/database"
	"errors"
	"time"
)

// User: Represents an API user
type User struct {
	Id        int64     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (user User) Create() error {
	stmt, err := db.Conn.Prepare("INSERT INTO users (name, email, nick, email, password) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return errors.New("cannot generate insert querry")
	}

	res, err := stmt.Exec(user.Name, user.Email)
	if err != nil {
		return errors.New("cannot insert data into database")
	}

	user.Id, err = res.LastInsertId()
	if err != nil {
		return errors.New("c~]annot search inserted id into database")
	}

	return nil
}
