package main

import (
	"fmt"

	"github.com/urfave/cli"
)

var shellCommand = cli.Command{
	Name:      "shell",
	Aliases:   []string{"sh"},
	Action:    shell,
	Usage:     "Run clico in shell mode.",
	UsageText: "clico shell",
}

func shell(c *cli.Context) {
	fmt.Printf("Just shelling...\n")
}
