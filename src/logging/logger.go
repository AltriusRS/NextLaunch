package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

type Logger struct {
	queue   chan string
	outputs []*Output
}

type Output struct {
	inner   *log.Logger
	Enabled bool
	TTY     bool
}

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

func (output *Output) Println(m string) {
	if output.Enabled {

		if output.TTY {
			output.inner.Println(m)
			return
		}

		// If the output isnt TTY, we need to strip the ANSI escape sequences
		// This is because the output is being piped to a file, and the ANSI escape sequences
		// will cause the file to be less readable
		output.inner.Println(regex.ReplaceAllString(m, ""))
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
		queue:   make(chan string, 100),
		outputs: []*Output{},
	}

	logger.outputs = append(logger.outputs, NewStdoutOutput(module))
	logger.outputs = append(logger.outputs, NewFileOutput(filepath.Join(cacheDir, "log.txt"), module))

	return logger
}

func (logger *Logger) Log(message string) {
	logger.queue <- message
}

func (logger *Logger) Logf(format string, v ...interface{}) {
	logger.Log(fmt.Sprintf(format, v...))
}

func (logger *Logger) Flush() {
	for {
		select {
		case message := <-logger.queue:
			for _, output := range logger.outputs {
				output.Println(message)
			}
		default:
			return
		}
	}
}

func (logger *Logger) run() {
	go func() {
		for {
			logger.Flush()
			time.Sleep(time.Millisecond * 10)
		}
	}()
}
