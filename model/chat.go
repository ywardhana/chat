package model

import "time"

const (
	ChatTimeFormat = "2006-01-02T15:04:05.000Z"
)

type Chat struct {
	message string
	time    time.Time
}

func NewChat(message string, time time.Time) *Chat {
	return &Chat{
		message: message,
		time:    time,
	}
}

func (c *Chat) GetMessage() string {
	return c.message
}

func (c *Chat) GetTime() string {
	return c.time.Format(ChatTimeFormat)
}
