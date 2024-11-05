package helper

import (
	db "devbookAPI/src/database"
	"devbookAPI/src/model"
	"errors"
	"fmt"
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

// SearchUser: Search users like recived string in database
func SearchUser(search string) ([]model.User, error) {
	search = fmt.Sprintf("%%%s%%", search)

	rows, err := db.Conn.Query(
		"SELECT id, name, nick, email, created_at FROM users WHERE (name LIKE ? OR nick LIKE ?)",
		search, search,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var user model.User

		if err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
