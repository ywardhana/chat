package http

import "strconv"

type createChatRequest struct {
	Message       string `json:"message"`
	InvoiceNumber string `json:"invoice_number"`
}

type ChatIndexParam struct {
	offset int
	limit  int
}

func NewChatIndexParam(conditions map[string][]string) (param ChatIndexParam, err error) {
	limit, offset, err := buildAttr(conditions)
	if err != nil {
		return
	}

	return ChatIndexParam{
		limit:  limit,
		offset: offset,
	}, nil
}

func (p ChatIndexParam) Limit() int {
	return p.limit
}

func (p ChatIndexParam) Offset() int {
	return p.offset
}

func buildAttr(conditions map[string][]string) (limit int, offset int, err error) {
	limit, err = buildLimit(conditions)
	if err != nil {
		return
	}
	offset, err = buildOffset(conditions)
	return
}

func buildLimit(conditions map[string][]string) (int, error) {
	return strconv.Atoi(conditions["limit"][0])
}

func buildOffset(conditions map[string][]string) (int, error) {
	return strconv.Atoi(conditions["offset"][0])
}
