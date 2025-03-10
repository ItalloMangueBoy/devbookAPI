package model

import (
	"encoding/json"
	"net/http"
	"time"
)

// Post: Represents one post
type Post struct {
	ID         int64     `json:"id,omitempty"`
	AuthorID   int64     `json:"author_id,omitempty"`
	AuthorNick string    `json:"author_nick,omitempty"`
	Content    string    `json:"content,omitempty"`
	Likes      int64     `json:"likes,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}

// Prepare: Format and validate the post data
func (post *Post) Prepare(r http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		return err
	}

	return post.format().validate()
}

func (post *Post) format() *Post {
	return post
}

func (post *Post) validate() error {
	return nil
}
