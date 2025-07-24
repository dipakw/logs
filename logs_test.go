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

func TestMust(t *testing.T) {
	logger := New(&Config{
		Allow: NONE,

		Outs: []*Out{
			{
				Target: os.Stdout,
				Color:  true,
			},
		},
	})

	logger.Must(INFO, DTAG, "Must information")
	logger.Must(WARN, DTAG, "Must warning")
	logger.Must(ERROR, DTAG, "Must error")

	fmt.Println()
}

func TestMustWithCustomTag(t *testing.T) {
	logger := New(&Config{
		Allow: NONE,

		Outs: []*Out{
			{
				Target: os.Stdout,
				Color:  true,
			},
		},
	})

	logger.Must(INFO, logger.Tag("infy | $d $t |"), "Must information custom tag")
	logger.Mustf(INFO, logger.Tag("infy | $d $t |"), "Must information custom tag and %s", "formatted")
	logger.Must(WARN, logger.Tag("wrny | $d $t |"), "Must warning custom tag")
	logger.Mustf(WARN, logger.Tag("wrny | $d $t |"), "Must warning custom tag and %s", "formatted")
	logger.Must(ERROR, logger.Tag("erry | $d $t |"), "Must error custom tag")
	logger.Mustf(ERROR, logger.Tag("erry | $d $t |"), "Must error custom tag and %s", "formatted")

	fmt.Println()
}

func TestMustWithNilTag(t *testing.T) {
	logger := New(&Config{
		Allow: NONE,

		Outs: []*Out{
			{
				Target: os.Stdout,
				Color:  true,
			},
		},
	})

	logger.Must(INFO, nil, "Must information nil tag")
	logger.Mustf(INFO, nil, "Must information nil tag and %s", "formatted")
	logger.Must(WARN, nil, "Must warning nil tag")
	logger.Mustf(WARN, nil, "Must warning nil tag and %s", "formatted")
	logger.Must(ERROR, nil, "Must error nil tag")
	logger.Mustf(ERROR, nil, "Must error nil tag and %s", "formatted")

	fmt.Println()
}
