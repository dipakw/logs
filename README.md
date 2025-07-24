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
		TagInf: "info : $t |", // Optional (default "inf | $d $t |) // $d = Current date (yyyy-mm-dd) $t = Current time (hh:mm:ss)
		TagWrn: "warn : $t |", // Optional (default "wrn | $d $t |)
		TagErr: "eror : $t |", // Optional (default "err | $d $t |)

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
