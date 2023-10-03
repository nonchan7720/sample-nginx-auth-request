package cmd

import (
	"log/slog"
	"os"
)

func Execute() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	log := slog.New(handler)
	slog.SetDefault(log)
	cmd := rootCommand()
	if err := cmd.Execute(); err != nil {
		log.Error(err.Error())
	}
}
