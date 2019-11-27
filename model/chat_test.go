package model_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ywardhana/chat/model"
)

func TestGetMessage(t *testing.T) {
	chatContent := "hallo"
	chat := model.NewChat(chatContent, time.Now())
	assert.Equal(t, chatContent, chat.GetMessage())
}

func TestGetTime(t *testing.T) {
	timeChatted := time.Now()
	chat := model.NewChat("hallo", timeChatted)
	assert.Equal(t, timeChatted.Format(model.ChatTimeFormat), chat.GetTime())
}
