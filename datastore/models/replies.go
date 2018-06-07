package models

import (
	"time"
)

type Reply struct {
	ID        uint
	UserID    uint
	CommentID uint
	CreatedAt time.Time
	Body      string
}
