package main

import "github.com/urfave/cli/v2"

var (
	pingFlag = &cli.BoolFlag{
		Name:    "show-ping-pong",
		Aliases: []string{"P"},
		Usage:   "print ping/pong",
	}
	headerFlag = &cli.StringSliceFlag{
		Name:    "header",
		Aliases: []string{"H"},
		Usage:   "add header(s) like curl",
	}
)
