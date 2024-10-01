package main

import (
	"context"
	"fmt"
	"os"
	"runtime"

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

	localos := runtime.GOOS
	localarch := runtime.GOARCH
	localshell := getShell()

	cmd := &cli.Command{
		Name:        "Clico",
		Usage:       "The CLI companion tool that brings the power of AI to your shell.",
		Description: clicoDescription,
		Copyright:   "2024 (c) António Maria Torre do Valle",
		UsageText:   "clico [global options] command [command options] \"prompt\"",
		Version:     "0.0.1",
		Commands:    []*cli.Command{pipeCommand, runCommand, askCommand},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "os",
				Usage:    "The operating system to use.",
				Value:    localos,
				Category: "Host information",
			},
			&cli.StringFlag{
				Name:     "arch",
				Usage:    "The architecture to use.",
				Value:    localarch,
				Category: "Host information",
			},
			&cli.StringFlag{
				Name:     "shell",
				Usage:    "The shell to use.",
				Value:    localshell,
				Category: "Host information",
			},
			&cli.StringFlag{
				Name:     "server",
				Usage:    "The server to connect to.",
				Value:    "http://localhost:11434",
				Sources:  cli.EnvVars("CLICO_SERVER"),
				Category: "LLM configuration",
			},
			&cli.StringFlag{
				Name:     "model",
				Usage:    "The model to use.",
				Value:    "llama3.1",
				Sources:  cli.EnvVars("CLICO_MODEL"),
				Category: "LLM configuration",
			},
			&cli.FloatFlag{
				Name:     "temperature",
				Usage:    "The temperature to use.",
				Value:    0.0,
				Sources:  cli.EnvVars("CLICO_TEMPERATURE"),
				Category: "LLM configuration",
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}

func getShell() string {
	shellVar := "SHELL"
	if runtime.GOOS == "windows" {
		shellVar = "COMSPEC"
	}

	shell := fmt.Sprintf("couldn't read from %s env var", shellVar)

	if os.Getenv(shellVar) != "" {
		shell = os.Getenv(shellVar)
	}

	return shell
}
