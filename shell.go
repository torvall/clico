package main

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

var shellCommand = &cli.Command{
	Name:      "shell",
	Aliases:   []string{"sh"},
	Action:    shell,
	Usage:     "Run clico in shell mode.",
	UsageText: "clico shell",
}

func shell(c context.Context, cmd *cli.Command) error {
	fmt.Printf("Just shelling...\n")
	return nil
}
