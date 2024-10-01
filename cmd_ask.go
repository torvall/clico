package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/urfave/cli/v3"
)

var askCommand = &cli.Command{
	Name:      "ask",
	Aliases:   []string{"a"},
	Action:    ask,
	Usage:     "Run clico in ask mode.",
	UsageText: "clico [global options] ask [options] \"prompt\"",
	Flags:     []cli.Flag{},
}

var askTemplate = `# You are Clico, the AI CLI companion

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

func ask(c context.Context, cmd *cli.Command) error {
	prompt := strings.Join(cmd.Args().Slice(), " ")

	// Bail out if there's no prompt.
	if len(prompt) == 0 {
		return fmt.Errorf("no prompt given")
	}

	localos := cmd.String("os")
	localarch := cmd.String("arch")
	localshell := cmd.String("shell")

	// Build prompt from template.
	promptStr := fmt.Sprintf(askTemplate, localos, localarch, localshell, prompt)

	// Query Ollama.
	outdata := queryOllama(promptStr)

	// Write data to stdout.
	fmt.Printf("%s\n", outdata)

	return nil
}
