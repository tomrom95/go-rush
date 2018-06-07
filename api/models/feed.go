package models

type FeedEventType uint

const (
	AddRushee       FeedEventType = 1
	CommentOnRushee FeedEventType = 2
	ReplyToComment  FeedEventType = 3
)

type RusheeFeedView struct {
	ID   uint   `json:id`
	Name string `json:name`
}

type UserFeedView struct {
	ID   uint   `json:id`
	Name string `json:name`
}

type FeedEvent struct {
	Rushee   RusheeFeedView `json:rushee`
	Actor    UserFeedView   `json:actor`
	Receiver UserFeedView   `json:receiver`
	Type     uint           `json:type`
}
