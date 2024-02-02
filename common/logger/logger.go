package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

type Logger interface {
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Warning(message string, args ...interface{})
	Error(message string, args ...interface{})
	Custom(prefix string, message string, args ...interface{})
}

type logger struct {
	logger *log.Logger
	level  int
}

func NewLogger(level int, prefix string) Logger {
	todo := color.New(color.FgHiWhite, color.Bold).SprintFunc()("ToDo")
	grenDot := color.New(color.FgGreen, color.Bold).SprintFunc()("•")
	app := color.New(color.FgHiWhite, color.Bold).SprintFunc()("APP®")

	rootPrefix := fmt.Sprintf("%s%s%s", todo, grenDot, app)

	if prefix != "" {
		__prefix := color.New(color.FgHiBlack, color.BgHiWhite, color.Bold).SprintFunc()(fmt.Sprintf(" %-20s ", prefix))
		prefix = fmt.Sprintf("%s %s ", rootPrefix, __prefix)
	}

	__log := log.New(os.Stdout, prefix, log.Ldate|log.Ltime)
	__log.SetPrefix(prefix)

	return &logger{
		logger: __log,
		level:  level,
	}
}

func (l *logger) Debug(message string, args ...interface{}) {
	prefix := color.New(color.FgHiBlack, color.BgHiWhite, color.Bold).SprintFunc()(" DEBUG ")
	l.logger.Println(prefix, fmt.Sprintf(message, args...))
	
}

func (l *logger) Info(message string, args ...interface{}) {
	prefix := color.New(color.FgHiWhite, color.BgCyan, color.Bold).SprintFunc()(" INFO  ")
	l.logger.Println(prefix, fmt.Sprintf(message, args...))

}

func (l *logger) Warning(message string, args ...interface{}) {
	prefix := color.New(color.FgHiBlack, color.BgHiYellow).Add(color.Bold).SprintFunc()(" WARN  ")
	l.logger.Println(prefix, fmt.Sprintf(message, args...))
	
}

func (l *logger) Error(message string, args ...interface{}) {
	prefix := color.New(color.FgHiWhite, color.BgHiRed).Add(color.Bold).SprintFunc()(" ERROR ")
	l.logger.Println(prefix, fmt.Sprintf(message, args...))
	
}

func (l *logger) Custom(prefix string, message string, args ...interface{}) {
	l.logger.Println(prefix, fmt.Sprintf(message, args...))
	
}
