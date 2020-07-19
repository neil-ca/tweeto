package models

import "time"

type RecordTweet struct {
	UserID string  `bson:"userid" json:"userid,omitempty"`
	Message string  `bson:"message" json:"message,omitempty"`
	Date time.Time  `bson:"date" json:"date,omitempty"`
}