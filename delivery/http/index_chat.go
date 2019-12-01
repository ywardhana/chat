package http

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ywardhana/goapi/response"
)

func (h *ChatHandler) IndexChat(w http.ResponseWriter, r *http.Request, param httprouter.Params) error {
	chatParam, err := NewChatIndexParam(r.URL.Query())
	if err != nil {
		return err
	}

	chats, err := h.usecase.IndexChat(chatParam)
	if err != nil {
		return err
	}
	total := h.usecase.CountChat()
	log.Println(chatParam.Offset())
	meta := response.MetaInfo{
		Limit:  chatParam.Limit(),
		Offset: chatParam.Offset(),
		Total:  total,
	}
	payload := serializerListChat(chats)
	response.OKWithMeta(w, payload, "", meta)
	return nil
}
