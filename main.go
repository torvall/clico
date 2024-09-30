// This is Clico, the CLI companion tool that brings the power of AI to your shell.
// It has 3 main modes of operation:
// - Run in shell mode: this is the default mode. You can run it with `clico` or `clico.exe`.
// - Run in pipe mode: in this mode, Clico can receive commands from stdin and retreive a command from the AI.
// - Print a command for a natural language query.
package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "clico"
	app.Usage = "CLI companion tool that brings the power of AI to your shell."
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		shellCommand,
		pipeCommand,
		queryCommand,
	}

	app.Run(os.Args)
}
