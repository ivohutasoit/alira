package message

import "time"

// Chat represents a single message which a client sent to a room
// (same meaning as a user send to a channel)
type Chat struct {
	Name      string    `json:"name" bson:"name"`
	Body      string    `json:"body" bson:"body"`
	Channel   string    `json:"channel" bson:"channel"`
	User      string    `json:"user" bson:"user"`
	Timestamp time.Time `json:"timestamp,omitempty" bson:"timestamp"`
}
