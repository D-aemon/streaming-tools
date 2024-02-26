package log

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
)

func init() {
	var lvl slog.LevelVar

	logLevel := os.Getenv("LOG_LEVEL")
	level, _ := strconv.Atoi(logLevel)

	lvl.Set(slog.Level(level))
	opts := slog.HandlerOptions{
		Level: &lvl,
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &opts)))

	slog.Info(fmt.Sprintf("Log Level: %s", slog.Level(level).String()))
}
