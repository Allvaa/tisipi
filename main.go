package main

import (
	"bufio"
	"os"
)

func main() {
	rw := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	rw.WriteString("What is your name? ")
	rw.Flush()

	str, _err := rw.ReadString('\n')

	if _err != nil {
		panic(_err)
	}

	rw.WriteString("Hello, " + str)
	rw.Flush()
}
