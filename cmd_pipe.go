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
	Usage:     "Transform piped in data",
	UsageText: "clico [global options] pipe [options] \"prompt\"",
	Flags:     []cli.Flag{},
}

var pipeTemplate = `# Data transformation request

Hello Clico, some information and a prompt follow, considering all the
context available to you, please change the data in the "stdin data"
section below according to the request in the "prompt" section.

Please respond only with the output of the transformation requested. If
no format was requested, use plain text, aligned with tabs if a listing.

## Host system information

This data is of the host system in question that the prompt refers to:

Operating system: ` + "`%s`" + `
Architecture: ` + "`%s`" + `
Shell: ` + "`%s`" + `

## Stdin data

This data was passed to the stdin, and is the input that the data transformation
request refers to:

` + "```" + `
%s
` + "```" + `

## Prompt

This is the data transformation operation requested by the user:

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
		return fmt.Errorf("no input data")
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
	promptStr := fmt.Sprintf(pipeTemplate, localos, localarch, localshell, string(indata), prompt)

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
