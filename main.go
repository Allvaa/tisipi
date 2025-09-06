package main

import (
	"fmt"
)

func main() {
	fmt.Printf("What is your name? ")
	var str string
	fmt.Scanf("%s", &str)

	fmt.Printf("Hello, %s\n", str)
}
