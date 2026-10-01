package main

import (
	"fmt"
	"os"

	"github.com/AliceFord/auto-sainsburys/internal/cli"
)

func main() {
	if err := cli.New().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
