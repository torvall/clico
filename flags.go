package main

import "github.com/urfave/cli"

var flagPrompt cli.StringFlag = cli.StringFlag{
	Name:  "prompt,p",
	Usage: "prompt for the query",
}
