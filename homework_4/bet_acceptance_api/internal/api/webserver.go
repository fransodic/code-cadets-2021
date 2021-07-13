package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

const betAcceptancePath = "/bet"

// WebServer Api server
type WebServer struct {
	router             *gin.Engine
	port               int
	readWriteTimeoutMs int
}

// NewServer returns new server instance
func NewServer(port, readWriteTimeoutMs int, ctrl Controller) *WebServer {
	server := &WebServer{
		router:             gin.Default(),
		port:               port,
		readWriteTimeoutMs: readWriteTimeoutMs,
	}
	server.registerRoutes(ctrl)
	return server
}

// Start on specified port and allow cancellation via context, if
// it crashes, cancel other goroutines via cancel function
func (w *WebServer) Start(ctx context.Context) {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", w.port),
		Handler:      w.router,
		ReadTimeout:  time.Duration(w.readWriteTimeoutMs) * time.Millisecond,
		WriteTimeout: time.Duration(w.readWriteTimeoutMs) * time.Millisecond,
	}
	errs := make(chan error)

	go func() {
		err := server.ListenAndServe()
		errs <- err
	}()

	log.Printf("Started http server, port: %s, host: %s\n", w.port, "127.0.0.1")

	select {
	case err := <-errs:
		log.Printf("An error occurred: %s", err.Error())
		return

	case <-ctx.Done():
		ctx, clear := context.WithTimeout(context.Background(), 1*time.Second)
		defer clear()

		// gracefully shutdown server
		err := server.Shutdown(ctx)

		if err != nil {
			log.Printf("An error occurred: %s", err.Error())
		}
		return
	}
}

// RegisterRoutes registers gin routes
func (w *WebServer) registerRoutes(ctrl Controller) {
	w.router.POST(betAcceptancePath, ctrl.AcceptBet())
}

// Controller handles api calls
type Controller interface {
	AcceptBet() gin.HandlerFunc
}
