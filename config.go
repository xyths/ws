package main

import "github.com/urfave/cli/v2"

var (
	pingFlag = &cli.BoolFlag{
		Name:    "show-ping-pong",
		Aliases: []string{"P"},
		Usage:   "print ping/pong",
	}
	urlFlag = &cli.StringFlag{
		Name:    "connect",
		Aliases: []string{"c"},
		Value:   "wss://echo.websocket.org",
		Usage:   "connect to a WebSocket server",
	}
	headerFlag = &cli.StringSliceFlag{
		Name:    "header",
		Aliases: []string{"H"},
	}
)
