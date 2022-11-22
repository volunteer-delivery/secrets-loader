package logger

import (
	"reflect"
	"strings"
)

type LogLevel int

const (
	None LogLevel = iota
	Error
	Debug
)

var logLevelMap = map[string]LogLevel{
	"none":  None,
	"error": Error,
	"debug": Debug,
}

var LogLevelNames = reflect.ValueOf(logLevelMap).MapKeys()

func ParseLogLevel(level string) (LogLevel, bool) {
	parsed, ok := logLevelMap[strings.ToLower(level)]
	return parsed, ok
}
