package server

import (
	"context"
	"math/rand"
	"net"
	"time"
)

var proverbs = []string{
	"Don't communicate by sharing memory, share memory by communicating.",
	"Concurrency is not parallelism.",
	"Channels orchestrate; mutexes serialize.",
	"The bigger the interface, the weaker the abstraction.",
	"Make the zero value useful.",
	"interface{} says nothing.",
	"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
	"A little copying is better than a little dependency.",
	"Syscall must always be guarded with build tags.",
	"Cgo must always be guarded with build tags.",
	"Cgo is not Go.",
	"With the unsafe package there are no guarantees.",
	"Clear is better than clever.",
	"Reflection is never clear.",
	"Errors are values.",
	"Don't just check errors, handle them gracefully.",
	"Design the architecture, name the components, document the details.",
	"Documentation is for users.",
	"Don't panic.",
}

const serverWisdomPause = time.Second * 3

func Listen(ctx context.Context, proto, socket string) error {

	listener, err := net.Listen(proto, socket)
	if err != nil {
		return err
	}
	defer func() { _ = listener.Close() }()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go handleProverb(ctx, conn)
	}

}

func handleProverb(ctx context.Context, c net.Conn) {
	defer func() { _ = c.Close() }()

	for {

		select {

		case <-ctx.Done():
			return

		case <-time.After(serverWisdomPause):

			rand.Seed(time.Now().UnixNano())

			p := proverbs[rand.Intn(len(proverbs))] + "\n"

			_, err := c.Write([]byte(p))
			if err != nil {
				return
			}

		}
	}
}
