package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"example.com/mod/api"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to run application", "error", err)
		os.Exit(1)
	}

	slog.Info("application stopped")
}

func run() error {
	apiHandler := api.NewHandler(make(map[string]string))

	srv := http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
		Addr:         ":8080",
		Handler:      apiHandler,
	}

	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
