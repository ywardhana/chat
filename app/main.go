package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/ywardhana/chat/app/system/middleware"
	"github.com/ywardhana/chat/app/system/server"
	chatHTTP "github.com/ywardhana/chat/delivery/http"
	"github.com/ywardhana/chat/repository"
	"github.com/ywardhana/chat/usecase"
)

func main() {
	chat := resolveChat()
	startServer(
		chat,
		// other handler goes here
	)
}

func resolveChat() server.Handler {
	chatRepository := repository.NewChatRepository()
	chatUsecase := usecase.NewChatUsecase(chatRepository)
	return chatHTTP.NewChatHandler(chatUsecase)

}

func startServer(handlers ...server.Handler) {
	mConfig := middleware.MiddlewareConfig{
		BasicUsername: "chat",
		BasicPassword: "chat",
	}
	m := middleware.NewMiddleware(mConfig)

	h := server.BuildServer(m, handlers...)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", 8001),
		Handler:      h,
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 300 * time.Second,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func(s *http.Server) {
		log.Printf("chat app is available at %s\n", s.Addr)
		if serr := s.ListenAndServe(); serr != http.ErrServerClosed {
			log.Fatal(serr)
		}
	}(s)

	<-sigChan

	log.Println("\nShutting down the app...")

	err := s.Shutdown(context.Background())
	if err != nil {
		log.Fatal("Something wrong when stopping server : ", err)
		return
	}

	log.Println("chat app is gracefully stopped")
}

type TestHandler struct {
}

func (h *TestHandler) Register(router *httprouter.Router, m *middleware.Middleware) {
	router.GET("/coba", m.AuthBasic(h.TestFunc))
}

func (h *TestHandler) TestFunc(w http.ResponseWriter, r *http.Request, param httprouter.Params) error {
	return nil
}
