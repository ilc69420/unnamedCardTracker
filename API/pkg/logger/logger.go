package logger

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"
)

type Logger struct {
	Base *slog.Logger
}

func New() Logger {
	handlerOpts := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}

	ip, hostName := GetHostInfo()

	baseLogger := slog.New(
		slog.NewJSONHandler(os.Stdout, handlerOpts),
	).With(
		"ip", ip,
		"host", hostName,
	)

	slog.SetDefault(baseLogger)

	return Logger{Base: baseLogger}
}

func (l *Logger) DB(msg, query, durationMs string) {
	l.Base.Info(msg, slog.Group("db",
		"query", "query",
		"durationMs", "durationMs",
	))
}

func (l *Logger) Controller(msg, controller string) {
	l.Base.Info(msg, slog.Group("controller",
		"name", "controllerName",
		"durationMs", "name",
	))
}

func GetHostInfo() (string, string) {
	cmd := exec.Command("tailscale", "status", "--self")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	parts := strings.Fields(string(out))
	if len(parts) > 2 {
		fmt.Println(parts[0], parts[1])
		return parts[0], parts[1]
	} else {
		fmt.Println("Could not get parts")
		panic(1)
	}
}
