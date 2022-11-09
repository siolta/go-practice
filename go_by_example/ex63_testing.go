// unit testing in go with the testing package
// run with: go test -v
package main

// testing code typically lives in the same package as the code it tests
import (
	"fmt"
	"testing"
)

// typically the code would be in intutils.go and the test file would be intutils_test.go
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// create a test by writing a func that begins with Test
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// t.Error* reports failures but doesn't halt the test, t.Fatal* stops the test
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// tests can be repetitive, so idiomatic to use a table-driven style
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{0, -1, -1},
	}

	// t.Run enables running subtests for each table entry
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
