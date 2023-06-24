// 5.
// Exercise 1.3: Experiment to measure the difference in running time
// between our potentially inefficient versions and the one that uses strings.Join.

package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
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

// BenchmarkEcho1
func BenchmarkEcho1(args []string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1(args)
	}
}

// BenchmarkEcho2
func BenchmarkEcho2(args []string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2(args)
	}
}

func main() {
	args := os.Args[1:]
	fmt.Println("echo1: ", echo1(args))
	fmt.Println("echo2: ", echo2(args))

	// run the benchmark tests
	fmt.Println("BenchmarkEcho1: ", testing.Benchmark(func(b *testing.B) {
		BenchmarkEcho1(args, b)
	}))

	fmt.Println("BenchmarkEcho2: ", testing.Benchmark(func(b *testing.B) {
		BenchmarkEcho2(args, b)
	}))
}
