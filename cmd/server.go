package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func newServer(router http.Handler) *Server {
	srv := &Server{
		apiServer: &http.Server{
			Handler:      router,
			ReadTimeout:  2 * time.Second,
			WriteTimeout: 5 * time.Second,
		},
	}
	return srv
}

type Server struct {
	apiServer *http.Server
}

func (srv Server) Serve(addr string) {
	srv.apiServer.Addr = addr
	go listenServer(srv.apiServer)
	waitForShutdown(srv.apiServer)
}

func listenServer(apiServer *http.Server) {
	err := apiServer.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatalf(err.Error())
	}
}

func waitForShutdown(apiServer *http.Server) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGINT,
		syscall.SIGTERM)
	_ = <-sig
	log.Printf("API server shutting down")
	_ = apiServer.Shutdown(context.Background())
	log.Printf("API server shutdown complete")
}
