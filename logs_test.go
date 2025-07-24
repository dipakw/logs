package logs

import (
	"fmt"
	"os"
	"testing"
)

func TestDefault(t *testing.T) {
	logger := New(&Config{
		Allow: ALL,

		Outs: []*Out{
			{
				Target: os.Stdout,
				Color:  true,
			},
		},
	})

	logger.Inf("Information")
	logger.Wrn("Warning")
	logger.Err("Error")

	fmt.Println()
}

func TestCustomTags(t *testing.T) {
	logger := New(&Config{
		Allow: ALL,

		TagInf: "info : $t |",
		TagWrn: "warn : $t |",
		TagErr: "eror : $t |",

		Outs: []*Out{
			{
				Target: os.Stdout,
				Color:  true,
			},
		},
	})

	logger.Inf("Information")
	logger.Wrn("Warning")
	logger.Err("Error")

	fmt.Println()
}
