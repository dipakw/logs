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
	})

	logger.Inf("Information")
	logger.Wrn("Warning")
	logger.Err("Error")

	fmt.Println()
}
