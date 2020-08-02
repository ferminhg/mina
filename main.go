package main

import (
	"os"

	"github.com/ferminhg/mina/cmd/cli"
)

func main() {
	cli.NewConsole(os.Stdout, os.Stderr, os.Exit).Run(os.Args)
}
