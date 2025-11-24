package logs

import (
	"fmt"
	"os"
)

func New(conf *Config) Logs {
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

	return &logger{
		cfg: conf,
	}
}

func (l *logger) Inf(a ...any) {
	l.write(Info, false, a...)
}

func (l *logger) Wrn(a ...any) {
	l.write(Warn, false, a...)
}

func (l *logger) Err(a ...any) {
	l.write(Error, false, a...)
}

func (l *logger) Inff(format string, a ...any) {
	l.writef(Info, false, format, a...)
}

func (l *logger) Wrnf(format string, a ...any) {
	l.writef(Warn, false, format, a...)
}

func (l *logger) Errf(format string, a ...any) {
	l.writef(Error, false, format, a...)
}

func (l *logger) Log(t Type, a ...any) {
	l.write(t, false, a...)
}

func (l *logger) Logf(t Type, format string, a ...any) {
	l.writef(t, false, format, a...)
}

func (l *logger) Write(a ...any) {
	l.write(Write, false, a...)
}

func (l *logger) Writef(format string, a ...any) {
	l.writef(Write, false, format, a...)
}

func (l *logger) Must(t Type, a ...any) {
	l.write(t, true, a...)
}

func (l *logger) Mustf(t Type, format string, a ...any) {
	l.writef(t, true, format, a...)
}

// ----- Log methods ----- //
func (l *Log) Deny() *Log {
	l.Allowed = false
	return l
}

func (l *Log) Allow() *Log {
	l.Allowed = true
	return l
}
