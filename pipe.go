package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

var pipeCommand = &cli.Command{
	Name:      "pipe",
	Aliases:   []string{"p"},
	Action:    pipe,
	Usage:     "Run clico in pipe mode.",
	UsageText: "clico pipe",
	Flags: []cli.Flag{
		flagPrompt,
	},
}

var pipeTemplate = `# You are Clico, the AI CLI companion

Hello and welcome to life! You are Clico, an AI powered CLI tool
that can be used to parse piped in data.

Any existing input is made available to you in the "stdin data" section
below. The user may refer to it in different ways or implicitly expect
you to know that's what they're referring to.

Please respond only with the result of the request.

Omit any prefixes or suffixes, don't use any markup.

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
	prompt := cmd.Value("prompt")

	// Get length of data in stdin.
	indatastat, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	indatalen := int(indatastat.Size())

	// Get any data in stdin.
	indata := make([]byte, indatalen)
	_, err = os.Stdin.Read(indata)
	if err != nil {
		panic(err)
	}

	// Build prompt from template.
	promptStr := fmt.Sprintf(pipeTemplate, indata, prompt)

	// Query Ollama.
	outdata := queryOllama(promptStr)

	// Write data to stdout.
	fmt.Printf("%s\n", outdata)

	return nil
}
