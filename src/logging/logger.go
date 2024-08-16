package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

type Level uint8

const (
	LevelNone Level = iota
	LevelFatal
	LevelError
	LevelWarning
	LevelInfo
	LevelDebug
)

type Logger struct {
	queue    chan Message
	outputs  []*Output
	panicked bool
}

type Message struct {
	Message string
	Level   Level
	Time    time.Time
}

type Output struct {
	inner   *log.Logger
	Enabled bool
	TTY     bool
}

func EnterTui() {
	MuteStdout = true
}

// MuteStdout Global variable to mute logging to stdout (used for showing the tui)
var MuteStdout = false

func NewFileOutput(source string, module string) *Output {

	file, err := os.OpenFile(source, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		log.Fatal(err)
	}

	return &Output{
		inner:   log.New(file, "["+module+"]: ", log.Lshortfile),
		Enabled: true,
		TTY:     false,
	}
}

func NewStdoutOutput(module string) *Output {
	return &Output{
		inner:   log.New(os.Stdout, "["+module+"]: ", log.Lshortfile),
		Enabled: true,
		TTY:     true,
	}
}

var regex = regexp.MustCompile(`\x1B(?:[@-Z\\-_]|\[[0-?]*[ -/]*[@-~])`)

func (output *Output) Println(m Message) {
	if output.Enabled {

		if !output.TTY {
			// If the output isn't a TTY, we need to strip the ANSI escape sequences
			// This is because the output is being piped to a file, and the ANSI escape sequences
			// will cause the file to be less readable
			m.Message = regex.ReplaceAllString(m.Message, "")
		}

		if output.TTY && MuteStdout {
			return
		}

		m.Message = m.Time.Format(time.RFC822Z) + " - " + m.Message

		switch m.Level {
		case LevelFatal:
			output.Enabled = false
			output.inner.Println("[Fatal] " + m.Message)
			return
		case LevelError:
			output.inner.Println("[Error] " + m.Message)
			return
		case LevelWarning:
			output.inner.Println("[Warning] " + m.Message)
			return
		case LevelInfo:
			output.inner.Println("[Info] " + m.Message)
			return
		case LevelDebug:
			output.inner.Println("[Debug] " + m.Message)
			return
		case LevelNone:
			output.inner.Println("[None] " + m.Message)
			return
		default:
			output.inner.Println("[None] " + m.Message)
			panic("unhandled default case in switch statement")
		}

	}
}

func (output *Output) Enable() {
	output.Enabled = true
}

func (output *Output) Disable() {
	output.Enabled = false
}

func (output *Output) IsEnabled() bool {
	return output.Enabled
}

func NewLogger(module string) *Logger {

	if module == "" {
		module = "NextLaunch"
	}

	var logger *Logger
	cacheDir, err := os.UserCacheDir()

	if err != nil {
		log.Fatal(err)
	}

	cacheDir = filepath.Join(cacheDir, "NextLaunch")

	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		err = os.Mkdir(cacheDir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	logger = &Logger{
		queue:    make(chan Message, 100),
		outputs:  []*Output{},
		panicked: false,
	}

	logger.outputs = append(logger.outputs, NewStdoutOutput(module))
	logger.outputs = append(logger.outputs, NewFileOutput(filepath.Join(cacheDir, "log.txt"), module))

	logger.run()

	return logger
}

func (logger *Logger) Log(message string) {
	logger.queue <- Message{
		Message: message,
		Level:   LevelInfo,
		Time:    time.Now(),
	}
}

func (logger *Logger) Logf(format string, v ...interface{}) {
	logger.Log(fmt.Sprintf(format, v...))
}

func (logger *Logger) Fatal(err error) {
	logger.queue <- Message{
		Message: err.Error(),
		Level:   LevelFatal,
		Time:    time.Now(),
	}
}

func (logger *Logger) Fatalf(format string, v ...interface{}) {
	logger.Fatal(fmt.Errorf(format, v...))
}

func (logger *Logger) Error(err error) {
	logger.queue <- Message{
		Message: err.Error(),
		Level:   LevelError,
		Time:    time.Now(),
	}
}

func (logger *Logger) Errorf(format string, v ...interface{}) {
	logger.Error(fmt.Errorf(format, v...))
}

func (logger *Logger) Warning(message string) {
	logger.queue <- Message{
		Message: message,
		Level:   LevelWarning,
		Time:    time.Now(),
	}
}

func (logger *Logger) Warningf(format string, v ...interface{}) {
	logger.Warning(fmt.Sprintf(format, v...))
}

func (logger *Logger) Info(message string) {
	logger.queue <- Message{
		Message: message,
		Level:   LevelInfo,
		Time:    time.Now(),
	}
}

func (logger *Logger) Infof(format string, v ...interface{}) {
	logger.Info(fmt.Sprintf(format, v...))
}

func (logger *Logger) Debug(message string) {
	logger.queue <- Message{
		Message: message,
		Level:   LevelDebug,
		Time:    time.Now(),
	}
}

func (logger *Logger) Debugf(format string, v ...interface{}) {
	logger.Debug(fmt.Sprintf(format, v...))
}

func (logger *Logger) Flush() {
	for {
		select {
		case message := <-logger.queue:
			atLeastOneOutput := false
			for _, output := range logger.outputs {
				output.Println(message)
				if output.Enabled {
					atLeastOneOutput = true
				}
			}

			if atLeastOneOutput {

			}
		default:
			return
		}
	}
}

func (logger *Logger) run() {
	go func() {
		for !logger.panicked {
			logger.Flush()
			time.Sleep(time.Millisecond * 10)
		}

		logger.Flush()
		println("Logger panicked and flushed")
	}()
	logger.Debug("Logger started")
}
