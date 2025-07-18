# Logs

A minimal logging library for Go.

## Example

```go
package main

import (
	"os"

	"github.com/dipakw/logs"
)

func main() {
	logger := logs.New(&logs.Config{
		Outs: []*logs.Out{
			{
				Target: os.Stdout,
				Color:  true,
			},
			{
				File:  "./output.log",
				Color: false,
			},
		},
	})

	logger.Inf("Some info")
	logger.Inff("Formated %s", "info")

	logger.Wrn("Some warning")
	logger.Wrnf("Formatted %s", "warn")

	logger.Err("Some error")
	logger.Errf("Formatted %s", "error")
}
```
