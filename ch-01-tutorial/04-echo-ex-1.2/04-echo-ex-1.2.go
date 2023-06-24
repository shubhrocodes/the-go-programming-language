// 4.
// Exercise 1.2: Modify the echo program to print the index and value of each of its arguments, one per line.

package main

import (
	"fmt"
	"os"
)

// echo3
func echo3(args []string) {
	for i, arg := range args {
		fmt.Printf("Index: %d, Value: %s\n", i, arg)
	}
}

func main() {
	args := os.Args[1:]
	echo3(args)
}
