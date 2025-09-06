package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

const PORT = 1550

func main() {
	listener, _err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))

	fmt.Printf("TCP server listening on port %d\n", PORT)

	if _err != nil {
		panic(_err)
	}

	for {
		conn, _err := listener.Accept()
		open := true

		if _err != nil {
			conn.Close()
			panic(_err)
		}

		reader := bufio.NewReader(conn)
		writer := bufio.NewWriter(conn)

		rw := bufio.NewReadWriter(reader, writer)

		for open {
			rw.WriteString("say hi: \n")
			rw.Flush()

			str, _err := rw.ReadString('\n')

			if _err != nil {
				panic(_err)
			}

			str = strings.TrimSpace(str)

			fmt.Printf("%s %s\n", conn.RemoteAddr().String(), str)

			switch str {
			case "hi":
				rw.WriteString("hello\n")
				rw.Flush()
			case "bye":
				conn.Close()
				open = false
			}
		}
	}
}
