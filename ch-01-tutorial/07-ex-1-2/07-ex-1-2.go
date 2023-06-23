package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("Index %d, Value %s\n", i, arg)
	}
}
