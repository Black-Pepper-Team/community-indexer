package main

import (
	"os"

	"github.com/black-pepper-team/community-indexer/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
