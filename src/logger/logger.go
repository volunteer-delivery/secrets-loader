package logger

import (
    "fmt"
)

type Logger struct {
    level LogLevel
}

func NewLogger(level LogLevel) *Logger {
    return &Logger{
        level: level,
    }
}

func (it *Logger) print(prefix string, message string, params []any) {
    output := fmt.Sprintf(message, params...)
    fmt.Println(fmt.Sprintf("%v --> %v", prefix, output))
}

func (it *Logger) canLog(level LogLevel) bool {
    return it.level >= level
}

func (it *Logger) printAtLevel(level LogLevel, prefix string, message string, params []any) {
    if it.canLog(level) {
        it.print(prefix, message, params)
    }
}

func (it *Logger) Data(message string, params ...any)  {
    fmt.Println(fmt.Sprintf(message, params...))
}

func (it *Logger) Error(message string, params ...any)  {
    it.printAtLevel(Error, "ERROR", message, params)
}

func (it *Logger) Debug(message string, params ...any)  {
    it.printAtLevel(Debug, "DEBUG", message, params)
}
