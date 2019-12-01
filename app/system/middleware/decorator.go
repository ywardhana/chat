package middleware

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ywardhana/chat/errormessage"
	"github.com/ywardhana/goapi/response"
)

type Decorator struct {
	successHandler    http.Handler
	failedAuthHandler http.Handler
	failedHandler     http.Handler
}

func NewDecorator() *Decorator {
	return &Decorator{
		successHandler:    http.HandlerFunc(HandlePassed),
		failedAuthHandler: http.HandlerFunc(HandleFailedAuth),
		failedHandler:     http.HandlerFunc(HandleFailed),
	}
}

func (d *Decorator) evaluate(auth Auth, r *http.Request) bool {
	return auth.Authenticate(r)
}

func (d *Decorator) ApplyDecorator(handler HandleWithError, auth Auth) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		if !d.evaluate(auth, r) {
			d.failedAuthHandler.ServeHTTP(w, r)
			return
		}
		err := handler(w, r, params)
		if err != nil {
			d.failedHandler.ServeHTTP(w, r)
		}
		d.successHandler.ServeHTTP(w, r)
	}
}

func HandleFailedAuth(w http.ResponseWriter, r *http.Request) {
	response.Error(w, errormessage.ErrorFailedAuth, http.StatusUnauthorized)
	log.Println(r.URL.Query())
	log.Output(1, errormessage.ErrorFailedAuth.Error()+"\n")
}

func HandlePassed(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Query())
	log.Println(r.URL.RequestURI())
	r.Context().Done()
}

func HandleFailed(w http.ResponseWriter, r *http.Request) {
	response.Error(w, errormessage.ErrorUnexpected, http.StatusUnprocessableEntity)
	log.Println(r.URL.Query())
}
