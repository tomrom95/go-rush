package models

type Event struct {
	ID    uint   `json:id`
	Name  string `json:name`
	Round uint   `json:round`
}

type Attendance struct {
	EventID  uint `json:event_id`
	RusheeID uint `json:rushee_id`
	UserID   uint `json:marked_by`
}
