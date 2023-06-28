// 7.
// Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countLines(f *os.File, counts map[string][]string) {
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		counts[line] = append(counts[line], f.Name())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading standard input: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	counts := make(map[string][]string)
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

	for line, files := range counts {
		if len(files) > 1 {
			fmt.Printf("%d\t%s\t%s\n", len(files), line, strings.Join(files, ", "))
		}
	}
}
