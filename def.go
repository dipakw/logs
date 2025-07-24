package logs

import "os"

const (
	TAG_INFO  = "inf | $d $t |"
	TAG_ERROR = "err | $d $t |"
	TAG_WARN  = "wrn | $d $t |"

	INFO  = 1
	WARN  = 2
	ERROR = 4
	ALL   = INFO | WARN | ERROR
	NONE  = 0
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
