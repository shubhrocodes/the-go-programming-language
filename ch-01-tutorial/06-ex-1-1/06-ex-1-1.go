package main

import (
	"fmt"
	"os"
)

func main() {
	s := os.Args[0] + " "
	for _, arg := range os.Args[1:] {
		s += arg + " "
	}
	fmt.Println(s)
}
