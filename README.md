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
		Allow: logs.All,

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
	logger.Inff("Formatted %s", "info")

	logger.Wrn("Some warning")
	logger.Wrnf("Formatted %s", "warning")

	logger.Err("Some error")
	logger.Errf("Formatted %s", "error")

	logger.Must(logs.Info, "Must information") // This gets logged regardless of the allow option.
	logger.Must(logs.Warn, "Must warning")
	logger.Mustf(logs.Error, "Must formatted %s", "error")
}
```

## Before Hook Example

The `Before` hook allows you to modify or filter logs before they are written:

```go
package main

import (
	"os"

	"github.com/dipakw/logs"
)

func main() {
	logger := logs.New(&logs.Config{
		Allow: logs.All,

		Outs: []*logs.Out{
			{
				Target: os.Stdout,
				Color:  true,
			},
		},

		Before: func(l *logs.Log) *logs.Log {
			// Filter out info logs
			if l.Type == logs.Info {
				return nil // Returning nil skips logging this entry
			}

			// Modify the tag for other log types
			l.Tag = "before | " + l.Tag
			return l
		},
	})

	logger.Inf("This info will be filtered out")
	logger.Wrn("This warning will have a modified tag")
	logger.Err("This error will have a modified tag")
}
```
