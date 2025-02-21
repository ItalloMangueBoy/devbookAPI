package followhelper

import (
	db "devbookAPI/src/database"
	"devbookAPI/src/model"
)

// GetUserFollowers: Get all followers of a user
func GetUserFollowers(userID int64) ([]model.User, error) {
	rows, err := db.Conn.Query(`
		SELECT u.id, u.name, u.nick, u.email, u.created_at FROM users u 
		INNER JOIN followers f ON  u.id = f.follower_id
		WHERE f.user_id = ?
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []model.User

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

		followers = append(followers, user)
	}

	return followers, nil
}

// FollowUser: Register a follower relationship in the database
func FollowUser(userID, followerID int64) error {
	stmt, err := db.Conn.Prepare("INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, followerID)
	if err != nil {
		return err
	}

	return nil
}

// UnfollowUser: Delete an follower relationship in the database
func UnfollowUser(userID, followerID int64) error {
	stmt, err := db.Conn.Prepare("DELETE FROM followers WHERE user_id = ? AND  follower_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, followerID)
	if err != nil {
		return err
	}

	return nil
}
