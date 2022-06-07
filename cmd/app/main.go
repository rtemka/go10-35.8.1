package main

import (
	"context"
	"log"
	"os"
	"proverbserver/pkg/client"
	"proverbserver/pkg/server"
	"time"
)

const proto = "tcp"

func main() {
	socket := os.Getenv("PROVERB_SERVER_SOCKET")
	if socket == "" {
		log.Fatal("$PROVERB_SERVER_SOCKET environment variable must be set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	go func() {
		err := server.Listen(ctx, proto, socket)
		if err != nil {
			log.Println(err)
		}
	}()

	err := client.Dial(ctx, proto, socket)
	if err != nil {
		log.Println(err)
	}
}
