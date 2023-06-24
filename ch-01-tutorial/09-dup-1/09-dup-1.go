// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter lines of text (type 'exit' to terminate): ")
	for input.Scan() {
		line := input.Text()
		if strings.ToLower(line) == "exit" {
			break
		}
		counts[line]++
	}
	fmt.Println("Lines that appear more than once: ")
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
