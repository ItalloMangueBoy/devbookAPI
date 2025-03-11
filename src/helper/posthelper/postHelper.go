package posthelper

import (
	db "devbookAPI/src/database"
	postsforms "devbookAPI/src/forms/postsForms"
	"devbookAPI/src/model"
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

// GetByID: search a post into database with given ID
func GetByID(ID int64) (model.Post, error) {
	var post model.Post

	err := db.Conn.
		QueryRow("SELECT posts.id, posts.content, posts.likes, posts.created_at, users.id, users.nick FROM posts INNER JOIN users ON posts.user_id = users.id WHERE posts.id = ?", ID).
		Scan(&post.ID, &post.Content, &post.Likes, &post.CreatedAt, &post.AuthorID, &post.AuthorNick)
	if err != nil {
		return model.Post{}, errors.New("error on server operation")
	}

	return post, nil
}
