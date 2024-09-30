package main

import (
	"fmt"

	"github.com/urfave/cli"
)

var queryCommand = cli.Command{
	Name:      "query",
	Aliases:   []string{"q"},
	Action:    query,
	Usage:     "Print a command for a natural language query.",
	UsageText: "clico query",
}

func query(c *cli.Context) {
	fmt.Printf("Querying for '%s'...\n", c.Args().First())
}
