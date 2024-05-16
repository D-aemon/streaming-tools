package logger

import (
	"backend/env"
	"fmt"
	"log/slog"
	"os"
	"strconv"
)

func InitSlog(config env.LogConfig) {
	var lvl slog.LevelVar

	level, _ := strconv.Atoi(config.Level)

	lvl.Set(slog.Level(level))
	opts := slog.HandlerOptions{
		Level: &lvl,
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &opts)))

	slog.Info(fmt.Sprintf("Log Level: %s", slog.Level(level).String()))
}
