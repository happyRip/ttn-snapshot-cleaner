package main

import (
	"os"

	"github.com/happyRip/ttn-snapshot-cleaner/cmd"
)

func main() {
	os.Exit(cmd.Execute())
}
