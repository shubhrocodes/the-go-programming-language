// 6.
// prints the count and text of lines that appear more than once in the input.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, counts map[string]int) {
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		counts[line]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading standard input: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Println("Please provide at least one input file.")
		return
	} else {
		for _, filename := range files {
			file, err := os.Open(filename)

			if err != nil {
				fmt.Fprintf(os.Stderr, "error opening file: %v\n", err)
				continue
			}
			countLines(file, counts)
			file.Close()
		}
	}

	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}
