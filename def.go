package logs

import "os"

type Type uint8

const (
	DefaultTagInfo  = "inf | $d $t | "
	DefaultTagError = "err | $d $t | "
	DefaultTagWarn  = "wrn | $d $t | "

	Info  Type = 1
	Warn  Type = 2
	Error Type = 4
	None  Type = 0
	All   Type = 1 | 2 | 4
)

var colors = map[Type]string{
	Error: "red",
	Warn:  "yellow",
}

type logger struct {
	cfg *Config
}

type Config struct {
	Allow  Type
	Outs   []*Out
	Tags   Tags
	Before func(l *Log) *Log
}

type Logs interface {
	Inf(a ...any)
	Wrn(a ...any)
	Err(a ...any)

	Inff(format string, a ...any)
	Wrnf(format string, a ...any)
	Errf(format string, a ...any)

	Must(t Type, a ...any)
	Mustf(t Type, format string, a ...any)
}

type Out struct {
	File   string
	Target *os.File
	Color  bool
}

type Log struct {
	Type    Type
	Message string
	Tag     string
	Allow   bool
	Must    bool
}

type Tags struct {
	Info  string
	Warn  string
	Error string
}
