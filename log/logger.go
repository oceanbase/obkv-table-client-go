/*-
 * #%L
 * OBKV Table Client Framework
 * %%
 * Copyright (C) 2021 OceanBase
 * %%
 * OBKV Table Client Framework is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *          http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 * #L%
 */

package log

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/sys/unix"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const ObkvTraceIdName = "_ObkvTraceIdName"

// Size specifies the maximum amount of data the writer will buffered before flushing. Defaults to 256 kB if unspecified.
const BufferSize = 4096

var (
	globalMutex         sync.Mutex
	defaultGlobalLogger *Logger
	SlowQueryThreshold  int64
)

type LogConfig struct {
	LogFileName        string // log file dir
	SingleFileMaxSize  int    // log file size（MB）
	MaxBackupFileSize  int    // Maximum number of old files to keep
	MaxAgeFileRem      int    // Maximum number of days to keep old files
	Compress           bool   // Whether to compress/archive old files
	SlowQueryThreshold int64  // Slow query threshold
}

type Level zapcore.Level

type Logger struct {
	l     *zap.Logger // zap ensure that zap.Logger is safe for concurrent use
	level Level
}

const (
	InfoLevel   Level = Level(zap.InfoLevel)   // 0, default level
	WarnLevel   Level = Level(zap.WarnLevel)   // 1
	ErrorLevel  Level = Level(zap.ErrorLevel)  // 2
	DPanicLevel Level = Level(zap.DPanicLevel) // 3, used in development log
	// PanicLevel logs a message, then panics
	PanicLevel Level = Level(zap.PanicLevel) // 4
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel Level = Level(zap.FatalLevel) // 5
	DebugLevel Level = Level(zap.DebugLevel) // -1
)

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

func AddInfo(logType string, traceId any, msg string) string {
	if len(logType) == 0 {
		logType = "Default"
	}
	var traceStr string
	if traceId == nil {
		traceStr = ""
	} else {
		traceStr = traceId.(string)
	}
	return "[" + logType + "] " + "[" + traceStr + "] " + msg
}

// Default
func Info(logType string, traceId any, msg string, fields ...Field) {
	defaultGlobalLogger.Info(AddInfo(logType, traceId, msg), fields...)
}

func Error(logType string, traceId any, msg string, fields ...Field) {
	defaultGlobalLogger.Error(AddInfo(logType, traceId, msg), fields...)
}

func Warn(logType string, traceId any, msg string, fields ...Field) {
	defaultGlobalLogger.Warn(AddInfo(logType, traceId, msg), fields...)
}

func DPanic(logType string, traceId any, msg string, fields ...Field) {
	defaultGlobalLogger.DPanic(AddInfo(logType, traceId, msg), fields...)
}

func Panic(logType string, traceId any, msg string, fields ...Field) {
	defaultGlobalLogger.Panic(AddInfo(logType, traceId, msg), fields...)
}

func Fatal(logType string, traceId any, msg string, fields ...Field) {
	defaultGlobalLogger.Fatal(AddInfo(logType, traceId, msg), fields...)
}

func Debug(logType string, traceId any, msg string, fields ...Field) {
	defaultGlobalLogger.Debug(AddInfo(logType, traceId, msg), fields...)
}

var (
	AddCaller     = zap.AddCaller
	WithCaller    = zap.WithCaller
	AddStacktrace = zap.AddStacktrace
	AddCallerSkip = zap.AddCallerSkip
)

func init() {
	globalMutex.Lock()
	defaultGlobalLogger = NewLogger(os.Stderr, InfoLevel, AddCaller())
	globalMutex.Unlock()
}

func InitLoggerWithConfig(cfg LogConfig) error {
	err := checkLoggerConfigValidity(cfg)
	if err != nil {
		return err
	}
	initDefaultLogger(cfg)
	return nil
}

func InitTraceId(ctx *context.Context) {
	if (*ctx).Value(ObkvTraceIdName) == nil {
		pid := uint64(unix.Getpid())
		// uniqueId uint64, pid uint64
		uniqueId := uuid.New().String()
		uniqueId = strings.ReplaceAll(uniqueId, "-", "")
		traceId := fmt.Sprintf("Y%s-%x", uniqueId, pid)
		*ctx = context.WithValue((*ctx), ObkvTraceIdName, traceId)
	}
}

func checkLoggerConfigValidity(cfg LogConfig) error {
	if cfg.LogFileName == "" {
		return errors.New("should set Log File Name in toml or client config")
	} else if cfg.SingleFileMaxSize <= 0 {
		return errors.New("Single File MaxSize is invalid")
	} else if cfg.SlowQueryThreshold <= 0 {
		return errors.New("Slow Query Threshold is invalid")
	} else if cfg.MaxAgeFileRem < 0 {
		return errors.New("Max Age File Remain is invalid")
	} else if cfg.MaxBackupFileSize < 0 {
		return errors.New("Max Backup File Size is invalid")
	}
	return nil
}

func initDefaultLogger(cfg LogConfig) {
	defFilePath := cfg.LogFileName + "obclient-table-go.log"
	SlowQueryThreshold = cfg.SlowQueryThreshold
	defLogWriter := getLogRotationWriter(defFilePath, cfg)
	globalMutex.Lock()
	defaultGlobalLogger = NewLogger(defLogWriter, InfoLevel, AddCaller())
	globalMutex.Unlock()
}

// rotation
func getLogRotationWriter(filePath string, cfg LogConfig) zapcore.WriteSyncer {
	asyncWrite := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   filePath,
			MaxSize:    cfg.SingleFileMaxSize,
			MaxBackups: cfg.MaxBackupFileSize,
			MaxAge:     cfg.MaxAgeFileRem,
			Compress:   cfg.Compress,
		}),
		//Size specifies the maximum amount of data the writer will buffered before flushing. Defaults to 256 kB if unspecified.
		Size: BufferSize, // async print buffer size
	}
	return asyncWrite
}

// no rotation
func getLogWriter(filePath string) zapcore.WriteSyncer {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil && os.IsNotExist(err) {
		file, _ = os.Create(filePath)
	}
	return zapcore.AddSync(file)
}

type Field = zap.Field

func (l *Logger) Debug(msg string, fields ...Field) {
	l.l.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.l.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.l.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.l.Error(msg, fields...)
}

func (l *Logger) DPanic(msg string, fields ...Field) {
	l.l.DPanic(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...Field) {
	l.l.Panic(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...Field) {
	l.l.Fatal(msg, fields...)
}

func (l *Logger) Sync() error {
	return l.l.Sync()
}

func Sync() error {
	if nil != defaultGlobalLogger {
		return defaultGlobalLogger.Sync()
	}
	return nil
}

func ResetDefaultLogger(l *Logger) {
	globalMutex.Lock()
	defaultGlobalLogger = l
	globalMutex.Unlock()
}

type Option = zap.Option

func NewLogger(writer io.Writer, level Level, opts ...Option) *Logger {
	if writer == nil {
		panic("the writer is nil")
	}
	var CustomEncoder = NewCustomEncoder()

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(CustomEncoder),
		zapcore.AddSync(writer),
		zapcore.Level(level),
	)

	logger := &Logger{
		l:     zap.New(core, opts...),
		level: level,
	}
	return logger
}

func NewCustomEncoder() zapcore.EncoderConfig {
	// time
	var customEncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("[" + "2006-01-02T15:04:05.000Z0700" + "]"))
	}
	// level
	customLevelEncoder := func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + level.CapitalString() + "]")
	}
	encoderConf := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level_name",
		MessageKey:     "msg",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     customEncodeTime,
		EncodeLevel:    customLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	return encoderConf
}

func MatchStr2LogLevel(level string) Level {
	switch level {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarnLevel
	case "error":
		return ErrorLevel
	case "dpanic":
		return DPanicLevel
	case "panic":
		return PanicLevel
	case "fatal":
		return FatalLevel
	default:
		return InfoLevel
	}
}
