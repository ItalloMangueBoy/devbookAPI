package followhelper

import (
	db "devbookAPI/src/database"
)

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
