package main

import (
	"fmt"
	"testing"
)

// Unit testing is an important part of writing principled Go programs

// The testing package provides the tools we need to write unit tests, and the go test command runs
// tests

// For the sake of demonstration, this code is in package main, but it could be any package

// Testing code typically lives in the same package as the code that it tests

// We'll be testing this simple implementation of an integer minimum
// Typically, the code we're testing would be in a source file named something like intutils.go, and
// the test file for it would then be named intutils_test.go
func IntMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// A test is created by writing a function with some name beginning with Test
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)

	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}

	// t.Error* will report test failures but continue executing the test
	// t.Fatal* will report test errors and stop the test immediately
}

// Writing test can be repetitive, so it's idiomatic to use a **table-driven style**, where test
// inputs and expected outputs are listed in a table, and a single loop walks over them and performs
// the test logic
func TestIntMinTableDrive(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	// t.Run enables running "subtests", one for each table entry
	// these are shown separately when executing go test -v
	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)

		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)

			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Benchmark tests typically go in _test.go files, and are named beginning with Benchmark

// The testing runner executes each benchmark function several times, increaseing b.N on each run
// until it collects a precise measurement

// Typically, the benchmark runs a function we're benchmarking in a loop b.N times
func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}

// Run tests in verbose mode with go test -v

// Then run all benchmarks in the current project with go test -bench=.
// (the -bench flag filters benchmark function names with a regex)
