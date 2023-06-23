package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func echo1(args []string) string {
	s := ""
	for _, arg := range args {
		s += arg + " "
	}
	return s
}

func echo2(args []string) string {
	return strings.Join(args, " ")
}

func BenchmarkEcho1(b *testing.B, args []string) {
	for i := 0; i < b.N; i++ {
		echo1(args)
	}
}

func BenchmarkEcho2(b *testing.B, args []string) {
	for i := 0; i < b.N; i++ {
		echo2(args)
	}
}

func main() {
	args := os.Args[1:]
	fmt.Println("Echo1:", echo1(args))
	fmt.Println("Echo2:", echo2(args))

	// Run the benchmark tests
	fmt.Println("Benchmark1:", testing.Benchmark(func(b *testing.B) {
		BenchmarkEcho1(b, args)
	}))
	fmt.Println("Benchmark2:", testing.Benchmark(func(b *testing.B) {
		BenchmarkEcho2(b, args)
	}))
}
