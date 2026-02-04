package logger

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type contextKey struct{}

// TraceInfo holds request tracing data
type TraceInfo struct {
	TraceID   string
	SpanID    string
	ParentID  string
	Component string
	Depth     int
}

// Config holds logger configuration
type Config struct {
	Level  string
	Pretty bool
}

var (
	globalLogger zerolog.Logger
	prettyMode   bool = false
)

// Init initializes the global logger
func Init(cfg Config) {
	if cfg.Level == "" {
		cfg.Level = "info"
	}

	lvl, _ := zerolog.ParseLevel(cfg.Level)
	zerolog.SetGlobalLevel(lvl)

	prettyMode = cfg.Pretty

	var output io.Writer = os.Stdout

	if cfg.Pretty {
		output = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "15:04:05",
			NoColor:    false,
			FormatLevel: levelFormatter,
			FormatMessage: func(i any) string {
				return fmt.Sprintf("  %s", i)
			},
			FormatFieldValue: func(i any) string {
				return fmt.Sprintf("\x1b[36m%v\x1b[0m", i)
			},
			FormatFieldName: func(i any) string {
				return fmt.Sprintf("\x1b[90m%s\x1b[0m", i)
			},
		}
	}

	globalLogger = zerolog.New(output).With().Timestamp().Logger()
	log.Logger = globalLogger
}

func levelFormatter(i any) string {
	level := i.(string)
	colors := map[string]string{
		"trace": "\x1b[90mTRC\x1b[0m",
		"debug": "\x1b[36mDBG\x1b[0m",
		"info":  "\x1b[32mINF\x1b[0m",
		"warn":  "\x1b[33mWRN\x1b[0m",
		"error": "\x1b[31mERR\x1b[0m",
		"fatal": "\x1b[35mFTL\x1b[0m",
		"panic": "\x1b[35mPNC\x1b[0m",
	}
	if c, ok := colors[level]; ok {
		return c
	}
	return level
}

// WithTrace adds trace info to context
func WithTrace(ctx context.Context, component string) context.Context {
	info := &TraceInfo{
		TraceID:   shortID(),
		SpanID:    shortID(),
		Component: component,
		Depth:     0,
	}
	return context.WithValue(ctx, contextKey{}, info)
}

// WithSpan creates a child span
func WithSpan(ctx context.Context, component string) context.Context {
	if parent, ok := ctx.Value(contextKey{}).(*TraceInfo); ok {
		info := &TraceInfo{
			TraceID:   parent.TraceID,
			SpanID:    shortID(),
			ParentID:  parent.SpanID,
			Component: component,
			Depth:     parent.Depth + 1,
		}
		return context.WithValue(ctx, contextKey{}, info)
	}
	return WithTrace(ctx, component)
}

// Global logging functions (use context for auto-tracing)

func Trace(ctx context.Context, msg string) {
	logWithLevel(ctx, zerolog.TraceLevel, msg)
}

func Debug(ctx context.Context, msg string) {
	logWithLevel(ctx, zerolog.DebugLevel, msg)
}

func Info(ctx context.Context, msg string) {
	logWithLevel(ctx, zerolog.InfoLevel, msg)
}

func Warn(ctx context.Context, msg string) {
	logWithLevel(ctx, zerolog.WarnLevel, msg)
}

func Error(ctx context.Context, msg string) {
	logWithLevel(ctx, zerolog.ErrorLevel, msg)
}

func Fatal(ctx context.Context, msg string) {
	logWithLevel(ctx, zerolog.FatalLevel, msg)
	os.Exit(1)
}

// Field helpers
type F map[string]any

// Log with fields
func Log(ctx context.Context, level string, msg string, fields F) {
	info, _ := ctx.Value(contextKey{}).(*TraceInfo)

	event := globalLogger.With()
	if info != nil {
		prefix := ""
		if prettyMode {
			for i := 0; i < info.Depth; i++ {
				prefix += "│  "
			}
			if info.Depth > 0 {
				prefix += "├─ "
			}
		}
		event = event.Str("trace", info.TraceID).Str("cmp", prefix+info.Component)
	}

	zlvl, _ := zerolog.ParseLevel(level)

	for k, v := range fields {
		event = event.Interface(k, v)
	}

	// Log with appropriate level
	l := event.Logger()
	switch zlvl {
	case zerolog.TraceLevel:
		l.Trace().Msg(msg)
	case zerolog.DebugLevel:
		l.Debug().Msg(msg)
	case zerolog.InfoLevel:
		l.Info().Msg(msg)
	case zerolog.WarnLevel:
		l.Warn().Msg(msg)
	case zerolog.ErrorLevel:
		l.Error().Msg(msg)
	case zerolog.FatalLevel:
		l.Fatal().Msg(msg)
	}
}

func logWithLevel(ctx context.Context, level zerolog.Level, msg string) {
	info, _ := ctx.Value(contextKey{}).(*TraceInfo)

	logger := globalLogger
	if info != nil {
		prefix := ""
		if prettyMode {
			for i := 0; i < info.Depth; i++ {
				prefix += "│  "
			}
			if info.Depth > 0 {
				prefix += "├─ "
			}
		}
		logger = globalLogger.With().Str("trace", info.TraceID).Str("cmp", prefix+info.Component).Logger()
	}

	logger.WithLevel(level).Msg(msg)
}

func shortID() string {
	return fmt.Sprintf("%x", time.Now().UnixNano())[:8]
}

// Request middleware helper
func RequestID(ctx context.Context) string {
	if info, ok := ctx.Value(contextKey{}).(*TraceInfo); ok {
		return info.TraceID
	}
	return "n/a"
}
