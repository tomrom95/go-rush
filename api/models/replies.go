package models

import (
	"time"
)

type Reply struct {
	ID        uint      `json:id`
	UserID    uint      `json:user_id`
	CommentID uint      `json:comment_id`
	CreatedAt time.Time `json:created_at`
	Body      string    `json:body`
}
