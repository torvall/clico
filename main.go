// This is Clico, the CLI companion tool that brings the power of AI to your shell.
// It has 3 main modes of operation:
// - Run in shell mode: this is the default mode. You can run it with `clico` or `clico.exe`.
// - Run in pipe mode: in this mode, Clico can receive commands from stdin and retreive a command from the AI.
// - Print a command for a natural language query.
package main

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:           "Clico",
		Description:    `Clico is a CLI companion tool that brings the power of AI to your shell.`,
		Aliases:        []string{},
		Authors:        []any{},
		Copyright:      "",
		Usage:          "",
		UsageText:      "",
		Version:        "",
		DefaultCommand: "pipe",
		Commands:       []*cli.Command{pipeCommand, shellCommand},
		Flags:          []cli.Flag{},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		panic(err)
	}
}
