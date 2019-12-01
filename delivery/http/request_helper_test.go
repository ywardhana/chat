package http_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const (
	BasicEncoded = "Y2hhdDpjaGF0"
)

type Request struct {
	method  string
	url     string
	payload interface{}
	auth    string
}

func NewRequest() *Request {
	return &Request{}
}

func (r *Request) Get(url string) *Request {
	return r.call("GET", url, nil)
}

func (r *Request) Post(url string, payload interface{}) *Request {
	return r.call("POST", url, payload)
}

func (r *Request) AsBasic() *Request {
	r.auth = "Basic " + BasicEncoded

	return r
}

func (r *Request) Build() *http.Request {
	req, _ := http.NewRequest(r.method, r.url, r.buildBody(r.payload))
	req.Header.Add("Content-Type", "application/json")

	if r.auth != "" {
		req.Header.Add("Authorization", r.auth)
	}

	return req
}

func (r *Request) buildBody(payload interface{}) io.Reader {
	j, _ := json.Marshal(payload)

	return bytes.NewReader(j)
}

func (r *Request) call(method string, url string, payload interface{}) *Request {
	r.method = method
	r.url = url
	r.payload = payload

	return r
}
