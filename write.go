package logs

import (
	"fmt"
	"strings"
	"time"

	"github.com/mgutz/ansi"
)

func (l *logger) _write(t Type, must bool, msg string) {
	if l.cfg.Outs == nil {
		return
	}

	var (
		tag   = ""
		allow = false
	)

	switch t {
	case Info:
		tag = strOr(DefaultTagInfo, l.cfg.Tags.Info)
		allow = (uint8(l.cfg.Allow) & uint8(Info)) > 0
	case Warn:
		tag = strOr(DefaultTagWarn, l.cfg.Tags.Warn)
		allow = (uint8(l.cfg.Allow) & uint8(Warn)) > 0
	case Error:
		tag = strOr(DefaultTagError, l.cfg.Tags.Error)
		allow = (uint8(l.cfg.Allow) & uint8(Error)) > 0
	default:
		tag = strOr(DefaultTagLog, l.cfg.Tags.Log)
		allow = l.cfg.Allow == All
	}

	log := &Log{
		Type:    t,
		Message: msg,
		Tag:     tag,
		Allow:   allow,
		Must:    must,
	}

	if l.cfg.Before != nil {
		log = l.cfg.Before(log)
	}

	if log == nil || (!must && !log.Allow) {
		return
	}

	tag = replaceTagVars(log.Tag)

	for _, out := range l.cfg.Outs {
		if out.Target == nil {
			continue
		}

		if out.Color && colors[t] != "" {
			out.Target.Write([]byte(ansi.Color(tag+log.Message, colors[t]) + "\n"))
		} else {
			out.Target.Write([]byte(tag + log.Message + "\n"))
		}
	}
}

func (l *logger) write(t Type, must bool, a ...any) {
	l._write(t, must, fmt.Sprint(a...))
}

func (l *logger) writef(t Type, must bool, format string, a ...any) {
	l._write(t, must, fmt.Sprintf(format, a...))
}

func strOr(a, b string) string {
	if a == "" {
		return b
	}

	return a
}

func replaceTagVars(t string) string {
	t = strings.ReplaceAll(t, "$d", time.Now().Format("2006-01-02"))
	t = strings.ReplaceAll(t, "$t", time.Now().Format("15:04:05.000"))
	return t
}
