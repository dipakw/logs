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

	logger.Must(Info, "Must information")
	logger.Mustf(Info, "Must formatted %s", "information")
	logger.Must(Warn, "Must warning")
	logger.Mustf(Warn, "Must formatted %s", "warning")
	logger.Must(Error, "Must error")
	logger.Mustf(Error, "Must formatted %s", "error")
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
}
