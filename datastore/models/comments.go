package models

import (
	"time"
)

type Comment struct {
	ID        uint
	RusheeID  uint
	UserID    uint
	CreatedAt time.Time
	Body      string
}
