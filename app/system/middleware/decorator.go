package middleware

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"github.com/ywardhana/chat/errormessage"
	"github.com/ywardhana/goapi/response"
)

type Decorator struct {
	successHandler    http.Handler
	failedAuthHandler http.Handler
	failedHandler     http.Handler
}

type websocketError struct {
	Message string `json:"message"`
}

func NewDecorator() *Decorator {
	return &Decorator{
		successHandler:    http.HandlerFunc(handlePassed),
		failedAuthHandler: http.HandlerFunc(handleFailedAuth),
		failedHandler:     http.HandlerFunc(handleFailed),
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
		if err := handler(w, r, params); err != nil {
			d.failedHandler.ServeHTTP(w, r)
		}
		d.successHandler.ServeHTTP(w, r)
	}
}

func (d *Decorator) ApplyWebsocketDecorator(handler ConnWithError) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		upgrader := websocket.Upgrader{}
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		if err := handler(conn); err != nil {
			log.Println(err)
			errMessage := websocketError{
				Message: errors.Wrap(err, "error").Error(),
			}
			conn.WriteJSON(errMessage)
		}
	}
}

func handleFailedAuth(w http.ResponseWriter, r *http.Request) {
	response.Error(w, errormessage.ErrorFailedAuth, http.StatusUnauthorized)
	log.Println(r.URL.Query())
	log.Output(1, errormessage.ErrorFailedAuth.Error()+"\n")
}

func handlePassed(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Query())
	log.Println(r.URL.RequestURI())
	r.Context().Done()
}

func handleFailed(w http.ResponseWriter, r *http.Request) {
	response.Error(w, errormessage.ErrorUnexpected, http.StatusUnprocessableEntity)
	log.Println(r.URL.Query())
}
