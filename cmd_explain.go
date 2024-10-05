package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v3"
)

var explainCommand = &cli.Command{
	Name:      "explain",
	Aliases:   []string{"e"},
	Action:    explain,
	Usage:     "Query the LLM using shell context information",
	UsageText: "clico [global options] explain [options] \"prompt\"",
	Flags:     []cli.Flag{},
}

var explainTemplate = `# Explanation request

Hello Clico, some information and a prompt follow, considering all the
context available to you, please try to address the prompt.

Any existing input is made available to you in the "stdin data" section
below. The prompt may refer to it in different ways or not at all. Whatever
explanation is requested, is about the contents of this data.

Please respond only with the explanation requested and be as concise as possible.

## Host system information

This data is of the host system in question that the prompt refers to:

Operating system: ` + "`%s`" + `
Architecture: ` + "`%s`" + `
Shell: ` + "`%s`" + `

## Stdin data

This data was passed to the stdin, and is the input that the question refers to:

` + "```" + `
%s
` + "```" + `

## Prompt

This is the user's actual question that you're answering:

` + "```" + `
%s
` + "```" + `
`

func explain(c context.Context, cmd *cli.Command) error {
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

	// Read input data, if any.
	strindata := "No input data available."
	if indatalen > 0 {
		// Get any data in stdin.
		indata := make([]byte, indatalen)
		_, err = os.Stdin.Read(indata)
		if err != nil {
			return err
		}
		strindata = string(indata)
	}

	localos := cmd.String("os")
	localarch := cmd.String("arch")
	localshell := cmd.String("shell")

	// Build prompt from template.
	promptStr := fmt.Sprintf(explainTemplate, localos, localarch, localshell, strindata, prompt)

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
