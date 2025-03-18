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
		QueryRow(`
			SELECT p.*, u.nick 
			FROM posts p 
			INNER JOIN users u ON p.user_id = u.id 
			WHERE p.id = ?
		`, ID).
		Scan(
			&post.ID,
			&post.AuthorID,
			&post.Content,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		)
	if err != nil {
		return model.Post{}, errors.New("error on server operation")
	}

	return post, nil
}

// GetTimeline: return the user timeline
func GetTimeline(ID int64) ([]model.Post, error) {
	// search posts in database
	rows, err := db.Conn.Query(`
			SELECT DISTINCT p.*, u.nick 
			FROM posts p 
			INNER JOIN users u ON p.user_id = u.id 
			INNER JOIN followers f ON u.id = f.user_id 
			WHERE u.id = ? OR f.follower_id = ?
			ORDER BY p.created_at DESC 
		`, ID, ID)

	if err != nil {
		return nil, errors.New("error on server operation")
	}
	defer rows.Close()

	// decode posts
	var timeline []model.Post

	for rows.Next() {
		var post model.Post
		if err := rows.Scan(
			&post.ID,
			&post.AuthorID,
			&post.Content,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, errors.New("error on server operation")
		}
		timeline = append(timeline, post)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.New("error on server operation")
	}

	// return success
	return timeline, nil
}

// ListUserPosts: return all posts from a user
func ListUserPosts(ID int64) ([]model.Post, error) {
	// execute query
	rows, err := db.Conn.Query(`
		SELECT p.*, u.nick 
		FROM posts p
		INNER JOIN users u ON u.id = p.user_id
		WHERE user_id = ?
	`, ID)
	if err != nil {
		return nil, errors.New("error on server operation")
	}
	defer rows.Close()

	// decode data
	var posts []model.Post

	for rows.Next() {
		var post model.Post

		if err := rows.Scan(
			&post.ID,
			&post.AuthorID,
			&post.Content,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, errors.New("error on server operation")
		}

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.New("error on server operation")
	}

	// return success
	return posts, nil
}

func Delete(ID int64) error {
	// remove post from database
	_, err := db.Conn.Exec("DELETE FROM posts WHERE id = ?", ID)
	if err != nil {
		return errors.New("error on server operation")
	}

	// return success
	return nil
}
