package main

import (
	"fmt"
	"os"
)

func test(arg string) {
	fmt.Printf("argument: %s\n", arg)
}

func main() {
	args := os.Args
	for _, arg := range args {
		test(arg)
	}
}
