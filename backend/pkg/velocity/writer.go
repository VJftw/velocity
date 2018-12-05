package velocity

import (
	"fmt"
)

type StreamWriter interface {
	Write(p []byte) (n int, err error)
	SetStatus(s string)
	Close()
}

// Emitter for forwarding bytes of output onwards
type Emitter interface {
	GetStreamWriter(streamName string) StreamWriter
}

type BlankEmitter struct {
}

func NewBlankEmitter() *BlankEmitter {
	return &BlankEmitter{}
}

func (w *BlankEmitter) GetStreamWriter(streamName string) StreamWriter {
	return &BlankWriter{}
}

type BlankWriter struct {
}

func (w BlankWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (w BlankWriter) SetStatus(s string) {}

func (w BlankWriter) Close() {}

const (
	ansiSuccess = "\x1b[1m\x1b[49m\x1b[32m"
	ansiWarn    = "\x1b[1m\x1b[49m\x1b[33m"
	ansiError   = "\x1b[1m\x1b[49m\x1b[31m"
	ansiInfo    = "\x1b[1m\x1b[49m\x1b[34m"
)

func colorFmt(ansiColor, format string) string {
	return fmt.Sprintf("%s%s\x1b[0m", ansiColor, format)
}
