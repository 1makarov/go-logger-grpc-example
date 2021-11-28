package types

import "time"

type LogItem struct {
	Entity    string    `bson:"entity"`
	Action    string    `bson:"action"`
	UserID    int64     `bson:"user_id"`
	Timestamp time.Time `bson:"timestamp"`
}
