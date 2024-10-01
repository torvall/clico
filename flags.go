package main

import "github.com/urfave/cli/v3"

var flagPrompt = &cli.StringFlag{
	Name:  "prompt,p",
	Usage: "prompt for the query",
}
