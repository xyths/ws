package main

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"strings"
)

func run(c *cli.Context) error {
	ping := c.Bool(pingFlag.Name)
	//log.Printf("ping: %v", ping)
	url := c.Args().Get(0)
	//log.Printf("connect to %s", url)
	var header http.Header
	headers := c.StringSlice(headerFlag.Name)
	if headers != nil {
		header = make(http.Header)
		for _, h := range headers {
			tokens := strings.Split(h, ":")
			if len(tokens) == 2 {
				header.Add(tokens[0], tokens[1])
			}
		}
		//log.Printf("headers: %v", headers)
		//log.Printf("header: %v", header)
	}

	conn, _, err := websocket.DefaultDialer.Dial(url, header)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("close connection error: %s", err)
		}
	}()

	done := make(chan struct{})

	go func(ctx context.Context) {
		defer close(done)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				msgType, message, err := conn.ReadMessage()
				if err != nil {
					log.Println("read:", err)
					return
				}
				switch msgType {
				case websocket.PingMessage:
					if ping {
						log.Println("received ping")
					}
				case websocket.TextMessage:
					log.Printf("%s", message)
				case websocket.PongMessage:
					if ping {
						log.Println("received pong")
					}
				}
			}
		}
	}(c.Context)

	for {
		select {
		case <-c.Context.Done():
			log.Println("ws stopped")
			return nil
		case _, ok := <-done:
			if !ok {
				log.Println("ws done closed")
			}
		}
	}
}
