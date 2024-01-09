package main

import (
	"os"

	"github.com/Distributed-Lab-Testing/example-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
