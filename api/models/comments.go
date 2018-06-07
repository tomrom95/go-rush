package models

import (
	"time"
)

type Comment struct {
	ID        uint      `json:id`
	RusheeID  uint      `json:rushee_id`
	UserID    uint      `json:user_id`
	CreatedAt time.Time `json:created_at`
	Body      string    `json:body`
}
