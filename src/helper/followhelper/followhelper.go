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