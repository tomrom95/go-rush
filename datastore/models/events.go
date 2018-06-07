package models

type Event struct {
	ID    uint
	Name  string
	Round uint
}

type Attendance struct {
	EventID  uint
	RusheeID uint
	UserID   uint // marked by
}
