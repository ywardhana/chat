package http

import "strconv"

const (
	DefaultLimit  = 20
	DefaultOffset = 0
)

type createChatRequest struct {
	Message string `json:"message"`
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

func buildLimit(conditions map[string][]string) (limit int, err error) {
	if len(conditions["limit"]) == 0 {
		limit = DefaultLimit
		return
	}
	limit, err = strconv.Atoi(conditions["limit"][0])
	return
}

func buildOffset(conditions map[string][]string) (offset int, err error) {
	if len(conditions["offset"]) == 0 {
		offset = DefaultOffset
		return
	}
	offset, err = strconv.Atoi(conditions["offset"][0])
	return
}
