package logs

import (
	"fmt"
	"os"
	"strings"

	"github.com/mgutz/ansi"
)

const (
	TAG_INFO  = "[..info...] "
	TAG_ERROR = "[..error..] "
	TAG_WARN  = "[..warn...] "

	TYPE_INFO  = 1
	TYPE_ERROR = 2
	TYPE_WARN  = 4
)

var colors = map[uint8]string{
	TYPE_ERROR: "red",
	TYPE_WARN:  "yellow",
}

type Logger struct {
	*Config
}

type Out struct {
	File   string
	Target *os.File
	Color  bool
}

type Config struct {
	NopOut bool
	Outs   []*Out
}

type Log interface {
	Inff(format string, a ...any)
	Wrnf(format string, a ...any)
	Errf(format string, a ...any)
	Inf(a ...any)
	Wrn(a ...any)
	Err(a ...any)
}

func New(c *Config) Log {
	if !c.NopOut {
		for _, out := range c.Outs {
			if out.File != "" {
				d, err := os.OpenFile(out.File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

				if err != nil {
					fmt.Println(err)
					os.Exit(1)

				}

				out.Target = d

				continue
			}
		}
	}

	return &Logger{c}
}

func (l *Logger) write(t uint8, s string) {
	if l.NopOut {
		return
	}

	for _, o := range l.Outs {
		if o.Target != nil {
			if o.Color && colors[t] != "" {
				o.Target.Write([]byte(ansi.Color(s, colors[t]) + "\n"))
			} else {
				o.Target.Write([]byte(s + "\n"))
			}
		}
	}
}

func (l *Logger) pp(t uint8, tag string, a ...any) {
	l.write(t, tag+strings.TrimSpace(fmt.Sprintln(a...)))
}

func (l *Logger) pf(t uint8, tag string, format string, a ...any) {
	l.write(t, fmt.Sprintf(tag+format, a...))
}

func (l *Logger) Inff(format string, a ...any) {
	l.pf(TYPE_INFO, TAG_INFO, format, a...)
}

func (l *Logger) Wrnf(format string, a ...any) {
	l.pf(TYPE_WARN, TAG_WARN, format, a...)
}

func (l *Logger) Errf(format string, a ...any) {
	l.pf(TYPE_ERROR, TAG_ERROR, format, a...)
}

func (l *Logger) Inf(a ...any) {
	l.pp(TYPE_INFO, TAG_INFO, a...)
}

func (l *Logger) Wrn(a ...any) {
	l.pp(TYPE_WARN, TAG_WARN, a...)
}

func (l *Logger) Err(a ...any) {
	l.pp(TYPE_ERROR, TAG_ERROR, a...)
}
