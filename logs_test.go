package logs

import (
	"fmt"
	"os"
	"testing"
)

func TestDefault(t *testing.T) {
	logger := New(&Config{
		Allow: All,

		Outs: []*Out{
			{
				Target: os.Stdout,
				Color:  true,
			},
		},

		Before: func(l *Log) *Log {
			return l
		},
	})

	logger.Inf("Information")
	logger.Inff("Formatted %s", "information")
	logger.Wrn("Warning")
	logger.Wrnf("Formatted %s", "warning")
	logger.Err("Error")
	logger.Errf("Formatted %s", "error")

	fmt.Println()
}

func TestMust(t *testing.T) {
	logger := New(&Config{
		Allow: None,

		Outs: []*Out{
			{
				Target: os.Stdout,
				Color:  true,
			},
		},
	})

	logger.Wrn("This should not be logged")
	logger.Must(Info, "Must information")
	logger.Mustf(Info, "Must formatted %s", "information")
	logger.Must(Warn, "Must warning")
	logger.Mustf(Warn, "Must formatted %s", "warning")
	logger.Must(Error, "Must error")
	logger.Mustf(Error, "Must formatted %s", "error")

	fmt.Println()
}

func TestLog(t *testing.T) {
	logger := New(&Config{
		Allow: All,
		Outs: []*Out{
			{
				Target: os.Stdout,
				Color:  true,
			},
		},
	})

	logger.Log(Info, "Log information")
	logger.Logf(Info, "Log formatted %s", "information")
	logger.Log(Warn, "Log warning")
	logger.Logf(Warn, "Log formatted %s", "warning")
	logger.Log(Error, "Log error")
	logger.Logf(Error, "Log formatted %s", "error")
	logger.Log(Write, "Log write")
	logger.Logf(Write, "Log formatted %s", "write")

	fmt.Println()
}

func TestWrite(t *testing.T) {
	logger := New(&Config{
		Allow: All,

		Outs: []*Out{
			{
				Target: os.Stdout,
				Color:  true,
			},
		},
	})

	logger.Write("Write information")
	logger.Writef("Write formatted %s", "information")

	fmt.Println()
}

func TestBefore(t *testing.T) {
	logger := New(&Config{
		Allow: All,

		Outs: []*Out{
			{
				Target: os.Stdout,
				Color:  true,
			},
		},

		Before: func(l *Log) *Log {
			if l.Type == Info {
				return nil
			}

			l.Tag = "before | " + l.Tag
			return l
		},
	})

	logger.Inf("Before information")
	logger.Inff("Before formatted %s", "information")
	logger.Wrn("Before warning")
	logger.Wrnf("Before formatted %s", "warning")
	logger.Err("Before error")
	logger.Errf("Before formatted %s", "error")

	fmt.Println()
}

func TestTypeFromString(t *testing.T) {
	logger := New(&Config{
		Allow: All,

		Outs: []*Out{
			{
				Target: os.Stdout,
				Color:  true,
			},
		},
	})

	logger.Log(T("info"), "Info log")
	logger.Log(T("warn"), "Warn log")
	logger.Log(T("error"), "Error log")
	logger.Log(T("invalid"), "Invalid log")

	fmt.Println()
}
