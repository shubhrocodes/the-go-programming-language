// 2.
// echo -  prints its command-line arguments

package main

import (
	"fmt"
	"os"
	"strings"
)

// echo1
func echo1(args []string) string {
	s := ""
	for _, arg := range args {
		s += arg + " "
	}
	return s
}

// echo2
func echo2(args []string) string {
	return strings.Join(args, " ")
}

func main() {
	args := os.Args[1:]
	fmt.Println(echo1(args))
	fmt.Println(echo2(args))
}
