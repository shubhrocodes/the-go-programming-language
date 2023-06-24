// 2.
// echo -  prints its command-line arguments

package main

import (
	"fmt"
	"os"
	"strings"
)

// 1st approach
func echo1(args []string) string {
	s := ""
	for _, arg := range args {
		s += arg + " "
	}
	return s
}

// 2nd approach
func echo2(args []string) string {
	return strings.Join(args, " ")
}

func main() {
	args := os.Args[1:]
	fmt.Println(echo1(args))
	fmt.Println(echo2(args))
}
