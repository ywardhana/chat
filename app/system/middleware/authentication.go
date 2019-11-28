package middleware

import "net/http"

type Auth interface {
	Authenticate(r *http.Request) bool
}

type BasicAuth struct {
	BasicUsername string
	BasicPassword string
}

func (b *BasicAuth) Authenticate(r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	return b.BasicUsername == username && b.BasicPassword == password && ok
}
