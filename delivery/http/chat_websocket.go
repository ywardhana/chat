package http

import (
	"github.com/gorilla/websocket"
)

func (h *ChatHandler) ChatWebsocket(c *websocket.Conn) error {
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			return err
		}
		h.usecase.CreateChat(string(message))
		chatParam, err := NewChatIndexParam(map[string][]string{})
		if err != nil {
			return err
		}

		chats, err := h.usecase.IndexChat(chatParam)
		if err != nil {
			return err
		}

		response := serializerListChat(chats)
		err = c.WriteJSON(response)
		if err != nil {
			return err
		}
	}
}
