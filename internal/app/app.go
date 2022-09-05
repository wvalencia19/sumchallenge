package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

func GracefullyShutdown(srv *http.Server) {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Panic("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	if <-ctx.Done(); true {
		log.Info("timeout of 5 seconds.")
	}
	log.Info("Server exiting")
}

func ParseLogLevel(level string) log.Level {
	l, err := log.ParseLevel(level)
	if err != nil {
		l = log.InfoLevel
		log.Errorf("setting log info %v", err)
	}
	return l
}
