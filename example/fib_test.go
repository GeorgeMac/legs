package fib

import (
	"fmt"
	"testing"

	"github.com/georgemac/legs"
)

type testCase struct {
	n        int // input
	expected int // expected result
}

func newCase(n, expected int) *testCase { return &testCase{n, expected} }

func (tt *testCase) Name() string {
	return fmt.Sprintf("Fib(%d) should equal %d", tt.n, tt.expected)
}

func (tt *testCase) Run(t *testing.T) {
	actual := Fib(tt.n)
	if actual != tt.expected {
		t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
	}
}

func TestFib(t *testing.T) {
	legs.Table([]legs.Case{
		newCase(1, 1),
		newCase(2, 1),
		newCase(3, 2),
		newCase(4, 3),
		newCase(5, 5),
		newCase(6, 8),
		newCase(7, 13),
	}).Run(t)
}
