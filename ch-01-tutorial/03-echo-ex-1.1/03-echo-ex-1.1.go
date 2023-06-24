// 3.
// Exercise 1.1: Modify the echo program to also print os.Args[0], the name of the command that invoked it.

package main

import (
	"fmt"
	"os"
)

// echo1
func echo1(args []string) string {
	s := os.Args[0] + " "
	for _, arg := range args {
		s += arg + " "
	}
	return s
}

func main() {
	args := os.Args[1:]
	fmt.Println(echo1(args))
}
