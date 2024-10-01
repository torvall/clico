package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v3"
)

var pipeCommand = &cli.Command{
	Name:      "pipe",
	Aliases:   []string{"p"},
	Action:    pipe,
	Usage:     "Transform piped in data.",
	UsageText: "clico [global options] pipe [options] \"prompt\"",
	Flags:     []cli.Flag{},
}

var pipeTemplate = `# You are Clico, the AI CLI companion

Hello and welcome to life! You are Clico, an AI powered CLI tool
that can be used to parse piped in data.

Any existing input is made available to you in the "stdin data" section
below. The user may refer to it in different ways or implicitly expect
you to know that's what they're referring to.

Please respond only with the result of the request.

Omit any prefixes or suffixes, don't use any markup.

## Host system information

Operating system: %s
Architecture: %s
Shell: %s

## Stdin data

` + "```" + `
%s
` + "```" + `

## Prompt

` + "```" + `
%s
` + "```" + `
`

func pipe(c context.Context, cmd *cli.Command) error {
	prompt := strings.Join(cmd.Args().Slice(), " ")

	// Bail out if there's no prompt.
	if len(prompt) == 0 {
		return fmt.Errorf("no prompt given")
	}

	// Get length of data in stdin.
	indatastat, err := os.Stdin.Stat()
	if err != nil {
		return err
	}
	indatalen := int(indatastat.Size())

	// Bail out if there's no data.
	if indatalen == 0 {
		return fmt.Errorf("no prompt given")
	}

	// Get any data in stdin.
	indata := make([]byte, indatalen)
	_, err = os.Stdin.Read(indata)
	if err != nil {
		return err
	}

	localos := cmd.String("os")
	localarch := cmd.String("arch")
	localshell := cmd.String("shell")

	// Build prompt from template.
	promptStr := fmt.Sprintf(pipeTemplate, localos, localarch, localshell, indata, prompt)

	server := cmd.String("server")
	model := cmd.String("model")
	temperature := cmd.Float("temperature")

	// Query Ollama.
	outdata, err := queryAPI(promptStr, server, model, temperature)
	if err != nil {
		return err
	}

	// Write data to stdout.
	fmt.Printf("%s\n", outdata)

	return nil
}
