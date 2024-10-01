package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v3"
)

var runCommand = &cli.Command{
	Name:      "run",
	Aliases:   []string{"r"},
	Action:    run,
	Usage:     "Run clico in run mode.",
	UsageText: "clico [global options] run [options] \"prompt\"",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "execute",
			Usage: "Execute commands instead of printing them.",
		},
	},
}

var runTemplate = `# You are Clico, the AI CLI companion

Hello and welcome to life! You are Clico, an AI powered CLI tool
that can be used to generate shell commands.

Please respond only with the command requested.

Omit any prefixes or suffixes, don't use any markup.

## Host system information

Operating system: %s
Architecture: %s
Shell: %s

## Prompt

` + "```" + `
%s
` + "```" + `
`

func run(c context.Context, cmd *cli.Command) error {
	prompt := strings.Join(cmd.Args().Slice(), " ")

	// Bail out if there's no prompt.
	if len(prompt) == 0 {
		return fmt.Errorf("no prompt given")
	}

	localos := cmd.String("os")
	localarch := cmd.String("arch")
	localshell := cmd.String("shell")

	// Build prompt from template.
	promptStr := fmt.Sprintf(runTemplate, localos, localarch, localshell, prompt)

	// Query Ollama.
	outdata := queryOllama(promptStr)

	execute := cmd.Bool("execute")

	if !execute {
		// Write data to stdout.
		fmt.Printf("%s\n", outdata)
	} else {
		// Run the command.
		cmdStr := strings.TrimSpace(outdata)
		fmt.Printf("executing %s...\n", cmdStr)
		cmd := exec.Command("sh", "-c", cmdStr)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	return nil
}
