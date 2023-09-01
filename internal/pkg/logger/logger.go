package logger

import (
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Level is a type alias for zapcore.Level
type Level = zapcore.Level

const (
	InfoLevel   Level = zapcore.InfoLevel  // 0, default level
	WarnLevel   Level = zapcore.WarnLevel  // 1
	ErrorLevel  Level = zapcore.ErrorLevel // 2
	DPanicLevel Level = zap.DPanicLevel    // 3, used in development log
	// PanicLevel logs a message, then panics
	PanicLevel Level = zapcore.PanicLevel // 4
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel Level = zapcore.FatalLevel // 5
	DebugLevel Level = zapcore.DebugLevel // -1
)

// Field is a type alias for zap.Field
type Field = zap.Field

var (
	Skip        = zap.Skip
	Binary      = zap.Binary
	Bool        = zap.Bool
	Boolp       = zap.Boolp
	ByteString  = zap.ByteString
	Complex128  = zap.Complex128
	Complex128p = zap.Complex128p
	Complex64   = zap.Complex64
	Complex64p  = zap.Complex64p
	Float64     = zap.Float64
	Float64p    = zap.Float64p
	Float32     = zap.Float32
	Float32p    = zap.Float32p
	Int         = zap.Int
	Intp        = zap.Intp
	Int64       = zap.Int64
	Int64p      = zap.Int64p
	Int32       = zap.Int32
	Int32p      = zap.Int32p
	Int16       = zap.Int16
	Int16p      = zap.Int16p
	Int8        = zap.Int8
	Int8p       = zap.Int8p
	String      = zap.String
	Stringp     = zap.Stringp
	Uint        = zap.Uint
	Uintp       = zap.Uintp
	Uint64      = zap.Uint64
	Uint64p     = zap.Uint64p
	Uint32      = zap.Uint32
	Uint32p     = zap.Uint32p
	Uint16      = zap.Uint16
	Uint16p     = zap.Uint16p
	Uint8       = zap.Uint8
	Uint8p      = zap.Uint8p
	Uintptr     = zap.Uintptr
	Uintptrp    = zap.Uintptrp
	Reflect     = zap.Reflect
	Namespace   = zap.Namespace
	Stringer    = zap.Stringer
	Time        = zap.Time
	Timep       = zap.Timep
	Stack       = zap.Stack
	StackSkip   = zap.StackSkip
	Duration    = zap.Duration
	Durationp   = zap.Durationp
	Any         = zap.Any
)

type Option = zap.Option

var (
	WithCaller    = zap.WithCaller
	WithStack     = zap.Stack
	AddStacktrace = zap.AddStacktrace
)

type Logger interface {
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	DPanic(msg string, fields ...Field)
	Panic(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
	With(fields ...Field) *zap.Logger
}

type RotateOptions struct {
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

type LevelEnablerFunc func(lvl Level) bool

type TeeOption struct {
	Filename string
	Ropt     RotateOptions
	Lef      LevelEnablerFunc
}

// AppLogger is an application error logs & warnings
type AppLogger struct {
	logger *zap.Logger
}

func NewAppLogger(logger *zap.Logger) *AppLogger {
	return &AppLogger{logger: logger}
}

func (l *AppLogger) Info(msg string, fields ...Field) {
	l.logger.Info(msg, fields...)
}

func (l *AppLogger) Warn(msg string, fields ...Field) {
	l.logger.Warn(msg, fields...)
}

func (l *AppLogger) Error(msg string, fields ...Field) {
	l.logger.Error(msg, fields...)
}

func (l *AppLogger) DPanic(msg string, fields ...Field) {
	l.logger.DPanic(msg, fields...)
}

func (l *AppLogger) Panic(msg string, fields ...Field) {
	l.logger.Panic(msg, fields...)
}

func (l *AppLogger) Fatal(msg string, fields ...Field) {
	l.logger.Fatal(msg, fields...)
}

func (l *AppLogger) Debug(msg string, fields ...Field) {
	l.logger.Debug(msg, fields...)
}

func (l *AppLogger) With(fields ...Field) *zap.Logger {
	return l.logger.With(fields...)
}

func (l *AppLogger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.Debug("OnStart hook executing: ",
			zap.String("callee", e.FunctionName),
			zap.String("caller", e.CallerName),
		)
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.Debug("OnStart hook failed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.Error(e.Err),
			)
		} else {
			l.Debug("OnStart hook executed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.String("runtime", e.Runtime.String()),
			)
		}
	case *fxevent.OnStopExecuting:
		l.Debug("OnStop hook executing: ",
			zap.String("callee", e.FunctionName),
			zap.String("caller", e.CallerName),
		)
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.Debug("OnStop hook failed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.Error(e.Err),
			)
		} else {
			l.Debug("OnStop hook executed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.String("runtime", e.Runtime.String()),
			)
		}
	case *fxevent.Supplied:
		l.Debug("supplied: ", zap.String("type", e.TypeName), zap.Error(e.Err))
	case *fxevent.Provided:
		for _, rtype := range e.OutputTypeNames {
			l.Debug("provided: ", zap.String("type", rtype))
		}
	case *fxevent.Decorated:
		for _, rtype := range e.OutputTypeNames {
			l.Debug("decorated: ",
				zap.String("decorator", e.DecoratorName),
				zap.String("type", rtype),
			)
		}
	case *fxevent.Invoking:
		l.Debug("invoking: ", zap.Any("func", e.FunctionName))
	case *fxevent.Started:
		if e.Err == nil {
			l.Debug("started")
		}
	case *fxevent.LoggerInitialized:
		if e.Err == nil {
			l.Debug("initialized: custom fxevent.Logger -> ", zap.Any("constructor", e.ConstructorName))
		}
	}
}

// Http Logger is an http request logs
// Http logger identify request by request id
type HttpLogger struct {
	logger *zap.Logger
}

func NewHttpLogger() *HttpLogger {
	return &HttpLogger{
		logger: zap.NewNop(),
	}
}

var _ Logger = (*HttpLogger)(nil)

func (l *HttpLogger) Info(msg string, fields ...Field) {
	l.logger.Info(msg, fields...)
}

func (l *HttpLogger) Warn(msg string, fields ...Field) {
	l.logger.Warn(msg, fields...)
}

func (l *HttpLogger) Error(msg string, fields ...Field) {
	l.logger.Error(msg, fields...)
}

func (l *HttpLogger) DPanic(msg string, fields ...Field) {
	l.logger.DPanic(msg, fields...)
}

func (l *HttpLogger) Panic(msg string, fields ...Field) {
	l.logger.Panic(msg, fields...)
}

func (l *HttpLogger) Fatal(msg string, fields ...Field) {
	l.logger.Fatal(msg, fields...)
}

func (l *HttpLogger) Debug(msg string, fields ...Field) {
	l.logger.Debug(msg, fields...)
}

func (l *HttpLogger) With(fields ...Field) *zap.Logger {
	return l.logger.With(fields...)
}
