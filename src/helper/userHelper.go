package helper

import (
	db "devbookAPI/src/database"
	"devbookAPI/src/model"
	"fmt"
)

// CreateUser: Insert one user in database
func CreateUser(user *model.User) error {
	stmt, err := db.Conn.Prepare("INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return err
	}

	user.Id, err = res.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

// SearchUsers: Search users like recived string in database
func SearchUsers(search string) ([]model.User, error) {
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

// GetUsers: Search users like recived string in database
func SearchUserById(id uint64) (user model.User, err error) {
	err = db.Conn.
		QueryRow("SELECT id, name, nick, email, created_at FROM users WHERE id = ?", id).
		Scan(&user.Id, &user.Name, &user.Nick, &user.Email, &user.CreatedAt)

	return
}

// UpdateUser: Updates one user in database
func UpdateUser(user *model.User) error {
	stmt, err := db.Conn.Prepare("UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(user.Name, user.Nick, user.Email, user.Id); err != nil {
		return err
	}

	return nil
}
