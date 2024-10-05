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
	Usage:     "Generate and optionally execute commands",
	UsageText: "clico [global options] run [options] \"prompt\"",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "execute",
			Usage: "Execute commands instead of printing them.",
		},
	},
}

var runTemplate = `# Command generation request

Hello Clico, here's a prompt for a shell command. Use the data in the
"host system information" below when generating the command that performs
the action described in the prompt.

Please respond only with the command requested. You can use multiple utilities
in the command to achieve the desired outcome.

Remember the command has to run on the host system described below.

## Host system information

This data is of the host system in question that the prompt refers to:

Operating system: ` + "`%s`" + `
Architecture: ` + "`%s`" + `
Shell: ` + "`%s`" + `

## Prompt

This is the command action we need to make a shell command of. Remember it
has to run on the host system described above:

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

	server := cmd.String("server")
	model := cmd.String("model")
	temperature := cmd.Float("temperature")

	// Query Ollama.
	outdata, err := queryAPI(promptStr, server, model, temperature)
	if err != nil {
		return err
	}

	execute := cmd.Bool("execute")

	if !execute {
		// Print the command.
		fmt.Printf("%s\n", outdata)
	} else {
		// Run the command.
		cmdStr := strings.TrimSpace(outdata)

		fmt.Printf("executing `%s`:\n", cmdStr)

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
