package logging

import (
	"os"

	"github.com/op/go-logging"
)

// Redactable is a redactable string type
type Redactable string

// Redacted is an method of Redactable to implement the Redactor interface
func (p Redactable) Redacted() interface{} {
	return logging.Redact(string(p))
}

// New returns a new Logger instance
func New() *logging.Logger {
	logger := logging.MustGetLogger("server")

	format := logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} >> %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)

	stderr := logging.NewLogBackend(os.Stderr, "", 0)
	stderrFormatted := logging.NewBackendFormatter(stderr, format)

	logging.SetBackend(stderrFormatted)

	return logger
}
