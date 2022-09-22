// Package logger provides
package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

// Level .
type Level int8

// Fields .
type Fields map[string]interface{}

const (
	// LevelDebug .
	LevelDebug Level = iota
	// LevelInfo .
	LevelInfo
	// LevelWarning .
	LevelWarning
	// LevelError .
	LevelError
	// LevelFatal .
	LevelFatal
	// LevelPanic .
	LevelPanic
)

// String .
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarning:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}
	return ""
}

// Logger .
type Logger struct {
	l       *log.Logger
	ctx     context.Context
	fields  Fields
	callers []string
	level   Level
}

// NewLogger .
func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	logger := log.New(w, prefix, flag)
	return &Logger{l: logger}
}

// clone .
func (l *Logger) clone() *Logger {
	newLogger := *l
	return &newLogger
}

// WithFields .
func (l *Logger) WithFields(f Fields) *Logger {
	logger := l.clone()
	if logger.fields == nil {
		logger.fields = make(Fields)
	}
	for k, v := range f {
		logger.fields[k] = v
	}
	return logger
}

// WithContext .
func (l *Logger) WithContext(ctx context.Context) *Logger {
	logger := l.clone()
	logger.ctx = ctx
	return logger
}

// WithCaller .
func (l *Logger) WithCaller(skip int) *Logger {
	logger := l.clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		logger.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}
	return logger
}

// WithCallersFrames .
func (l *Logger) WithCallersFrames() *Logger {
	callers := []string{}

	maxCallerDepth := 25
	minCallerDepth := 1

	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		s := fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function)
		callers = append(callers, s)
		if !more {
			break
		}
	}

	logger := l.clone()
	logger.callers = callers
	return logger
}

// WithTrace .
func (l *Logger) WithTrace() *Logger {
	ginCtx, ok := l.ctx.(*gin.Context)
	if ok {
		l.WithFields(Fields{
			"trace_id": ginCtx.MustGet("X-Trace-ID"),
			"span_id":  ginCtx.MustGet("X-Span-ID"),
		})
	}
	return l
}

// WithLevel .
func (l *Logger) WithLevel(level Level) *Logger {
	logger := l.clone()
	logger.level = level
	return logger
}

// JSONFormat .
func (l *Logger) JSONFormat(level Level, message string) map[string]interface{} {
	data := make(Fields, len(l.fields)+4)
	data["level"] = level.String()
	data["time"] = time.Now().UnixNano()
	data["message"] = message
	data["callers"] = l.callers
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			data[k] = v
		}
	}
	return data
}

// Output .
func (l *Logger) Output(message string) {
	b, _ := json.Marshal(l.JSONFormat(l.level, message))
	content := string(b)
	switch l.level {
	case LevelDebug:
		l.l.Print(content)
	case LevelInfo:
		l.l.Print(content)
	case LevelWarning:
		l.l.Print(content)
	case LevelError:
		l.l.Print(content)
	case LevelFatal:
		l.l.Fatal(content)
	case LevelPanic:
		l.l.Panic(content)
	}
}

// Info .
func (l *Logger) Info(ctx context.Context, v ...interface{}) {
	l = l.WithLevel(LevelInfo).WithContext(ctx).WithTrace()
	l.Output(fmt.Sprint(v...))
}

// Infof .
func (l *Logger) Infof(ctx context.Context, format string, v ...interface{}) {
	l = l.WithLevel(LevelInfo).WithContext(ctx).WithTrace()
	l.Output(fmt.Sprintf(format, v...))
}

// Error .
func (l *Logger) Error(ctx context.Context, v ...interface{}) {
	l = l.WithLevel(LevelError).WithContext(ctx).WithTrace()
	l.Output(fmt.Sprint(v...))
}

// Errorf .
func (l *Logger) Errorf(ctx context.Context, format string, v ...interface{}) {
	l = l.WithLevel(LevelError).WithContext(ctx).WithTrace()
	l.Output(fmt.Sprintf(format, v...))
}
