package posthelper

import (
	db "devbookAPI/src/database"
	postsforms "devbookAPI/src/forms/postsForms"
	"errors"
)

// Create: insert a new post in deatabase
func Create(data *postsforms.Create) error {
	// prepere query
	stmt, err := db.Conn.Prepare("INSERT INTO posts (user_id, content) VALUES (?, ?)")
	if err != nil {
		return errors.New("error on server operation")
	}
	defer stmt.Close()

	// execute query
	if _, err := stmt.Exec(data.AuthorID, data.Content); err != nil {
		return errors.New("error on server operation")
	}

	// return success
	return nil
}
