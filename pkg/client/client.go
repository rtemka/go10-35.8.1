package client

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"time"
)

const timeout = time.Second * 5

func Dial(ctx context.Context, proto, socket string) error {
	childCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var d net.Dialer
	conn, err := d.DialContext(childCtx, proto, socket)
	if err != nil {
		return err
	}
	defer func() { _ = conn.Close() }()

	r := bufio.NewReader(conn)

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			p, err := r.ReadBytes('\n')
			if err != nil {
				if err != io.EOF {
					return err
				}
				return nil
			}

			fmt.Printf("received: %s\n", string(p))
		}
	}
}
