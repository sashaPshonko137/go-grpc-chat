package message

import "time"

type Message struct {
	MessageId int32
	UserId int32
	Text string
	Time time.Time
}

type MessageInfo struct {
	UserId int32
	ChatId int32
	Text string
	Time time.Time
}