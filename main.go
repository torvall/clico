package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	clicoDescription := `This is Clico, a CLI COmpanion tool that enables you to use AI
to manipulate output, recommend commands or query an LLM using
contextual data from your shell history and output.

It's designed to allow complex operations in data to be described
using natural language, serve as a copilot in your shell to help
you develop or debug, or as a troubleshooting tool. Can also be
used in automation.

⚠️ WARNING: Clico can automatically execute commands generated
by the LLM. Use with caution.`

	cmd := &cli.Command{
		Name:        "Clico",
		Usage:       "The CLI companion tool that brings the power of AI to your shell.",
		Description: clicoDescription,
		Copyright:   "2024 (c) António Maria Torre do Valle",
		UsageText:   "clico [global options] command [command options] \"prompt\"",
		Version:     "0.0.1",
		Commands:    []*cli.Command{pipeCommand, shellCommand},
		Flags:       []cli.Flag{},
		// DefaultCommand: "pipe",
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}
