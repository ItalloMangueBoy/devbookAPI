package model

import "time"

type Postes struct {
	ID        uint64    `json:"id,omitempty"`
	UserID    uint64    `json:"user_id,omitempty"`
	UserNick  string    `json:"user_nick,omitempty"`
	Content   string    `json:"content,omitempty"`
	Likes     uint64    `json:"likes,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
