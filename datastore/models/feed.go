package models

type FeedEventType uint

const (
	AddRushee       FeedEventType = 1
	CommentOnRushee FeedEventType = 2
	ReplyToComment  FeedEventType = 3
)

type FeedEvent struct {
	RusheeID   uint
	Rushee     Rushee
	ActorID    uint
	Actor      User
	ReceiverID *uint
	Receiver   *User
	Type       FeedEventType
}
