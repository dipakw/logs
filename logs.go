package logs

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mgutz/ansi"
)

const (
	TAG_INFO  = "inf | $d $t |"
	TAG_ERROR = "err | $d $t |"
	TAG_WARN  = "wrn | $d $t |"

	INFO  = 1
	WARN  = 2
	ERROR = 4
	ALL   = INFO | WARN | ERROR
)

var colors = map[uint8]string{
	ERROR: "red",
	WARN:  "yellow",
}

type Logger struct {
	conf *Config
}

type Out struct {
	File   string
	Target *os.File
	Color  bool
}

type Config struct {
	NopOut bool
	Allow  uint8
	Outs   []*Out
	TagInf string
	TagWrn string
	TagErr string
}

type Log interface {
	Inff(format string, a ...any)
	Wrnf(format string, a ...any)
	Errf(format string, a ...any)
	Inf(a ...any)
	Wrn(a ...any)
	Err(a ...any)
}

func New(conf *Config) Log {
	for _, out := range conf.Outs {
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

	return &Logger{
		conf: conf,
	}
}

func (l *Logger) write(t uint8, s string) {
	if l.conf.NopOut || (l.conf.Allow&t) == 0 {
		return
	}

	for _, o := range l.conf.Outs {
		if o.Target != nil {
			if o.Color && colors[t] != "" {
				o.Target.Write([]byte(ansi.Color(s, colors[t]) + "\n"))
			} else {
				o.Target.Write([]byte(s + "\n"))
			}
		}
	}
}

func (l *Logger) getTag(t uint8) string {
	tag := ""

	switch t {
	case INFO:
		tag = strOr(l.conf.TagInf, TAG_INFO)
	case WARN:
		tag = strOr(l.conf.TagWrn, TAG_WARN)
	case ERROR:
		tag = strOr(l.conf.TagErr, TAG_ERROR)
	}

	tag = strings.ReplaceAll(tag, "$d", time.Now().Format("2006-01-02"))
	tag = strings.ReplaceAll(tag, "$t", time.Now().Format("15:04:05"))

	return tag + " "
}

func (l *Logger) pp(t uint8, a ...any) {
	tag := l.getTag(t)
	l.write(t, tag+strings.TrimSpace(fmt.Sprintln(a...)))
}

func (l *Logger) pf(t uint8, format string, a ...any) {
	tag := l.getTag(t)
	l.write(t, fmt.Sprintf(tag+format, a...))
}

func (l *Logger) Inff(format string, a ...any) {
	l.pf(INFO, format, a...)
}

func (l *Logger) Wrnf(format string, a ...any) {
	l.pf(WARN, format, a...)
}

func (l *Logger) Errf(format string, a ...any) {
	l.pf(ERROR, format, a...)
}

func (l *Logger) Inf(a ...any) {
	l.pp(INFO, a...)
}

func (l *Logger) Wrn(a ...any) {
	l.pp(WARN, a...)
}

func (l *Logger) Err(a ...any) {
	l.pp(ERROR, a...)
}

func strOr(s string, d string) string {
	if s == "" {
		return d
	}

	return s
}
