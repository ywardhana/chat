package http

import (
	"encoding/json"
	"net/http"

	"github.com/ywardhana/goapi/response"

	"github.com/julienschmidt/httprouter"
)

func (h *ChatHandler) CreateChat(w http.ResponseWriter, r *http.Request, param httprouter.Params) error {
	var chatRequest createChatRequest

	if err := json.NewDecoder(r.Body).Decode(&chatRequest); err != nil {
		return err
	}
	chat := h.usecase.CreateChat(chatRequest.Message)

	responseData := serializerChat(chat)

	response.OK(w, responseData, "")
	return nil
}
