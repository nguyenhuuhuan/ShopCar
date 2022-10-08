package logger

import (
	"context"
	"errors"
)

// A global variable so that log functions can be directly accessed
var log LoggerV2

func init() {
	err := NewLogger(Configuration{
		EnableConsole:     true,
		ConsoleLevel:      Debug,
		ConsoleJSONFormat: true,
		EnableFile:        true,
	}, InstanceZapLogger)

	if err != nil {
		log.Fatalf(err, "Could not instantiate log %v", err)
	}
}

//Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

const (
	//Debug has verbose message
	Debug = "debug"
	//Info is default log level
	Info = "info"
	//Warn is for logging messages about possible issues
	Warn = "warn"
	//Error is for logging errors
	Error = "error"
	//Fatal is for logging fatal messages. The sytem shutsdown after logging the message.
	Fatal = "fatal"
)

const (
	InstanceZapLogger int = iota
	InstanceLogrusLogger
)

var (
	errInvalidLoggerInstance = errors.New("Invalid logger instance")
)

//LoggerV2 is our contract for the logger
type LoggerV2 interface {
	Debugf(format string, args ...interface{})

	Infof(format string, args ...interface{})
	Infow(args ...interface{})

	Warnf(format string, args ...interface{})

	Errorf(format string, args ...interface{})

	Fatalf(err error, format string, args ...interface{})

	Panicf(format string, args ...interface{})

	Report(err error, format string, args ...interface{})

	WithFields(keyValues Fields) LoggerV2
}

// Configuration stores the config for the logger
// For some loggers there can only be one level across writers, for such the level of Console is picked by default
type Configuration struct {
	EnableConsole     bool
	ConsoleJSONFormat bool
	ConsoleLevel      string
	EnableFile        bool
	FileJSONFormat    bool
}

//NewLogger returns an instance of logger
func NewLogger(config Configuration, loggerInstance int) error {
	switch loggerInstance {
	case InstanceZapLogger:
		logger, err := newZapLogger(config)
		if err != nil {
			return err
		}
		log = logger
		return nil
	default:
		return errInvalidLoggerInstance
	}
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Infow(args ...interface{}) {
	log.Infow(args)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Fatalf(err error, format string, args ...interface{}) {
	log.Fatalf(err, format, args...)
}

func Report(err error, format string, args ...interface{}) {
	log.Report(err, format, args...)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func WithFields(keyValues Fields) LoggerV2 {
	return log.WithFields(keyValues)
}

func Log() LoggerV2 {
	return log
}

func Context(ctx context.Context) LoggerV2 {
	if ctx != nil {
		if ctxRqID, ok := ctx.Value(RqIDCtxKey).(string); ok {
			return log.WithFields(Fields{
				"rqID": ctxRqID,
			})
		}
	}
	return log
}
