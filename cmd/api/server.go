package api

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (a *app) StartServer() error {
	srv := &http.Server{
		Addr:     a.config.Addr,
		Handler:  a.routes(),
		ErrorLog: slog.NewLogLogger(a.logger.Handler(), slog.LevelError),
	}
	a.logger.Info("starting server", "addr", srv.Addr, "env", a.config.Env)

	shutdownErr := make(chan error)
	go func() {
		done := make(chan os.Signal)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-done
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		shutdownErr <- srv.Shutdown(ctx)
	}()

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	err := <-shutdownErr
	if err != nil {
		a.logger.Error(err.Error())
	}
	a.logger.Info("Server stopped")

	return err

}
