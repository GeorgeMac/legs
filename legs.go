package legs

import "testing"

// Runner can run a test given a pointer to a testing.T
type Runner interface {
	Run(t *testing.T)
}

// RunnerFunc is a helper type which implements
// Runner and delegates to the decorated function
type RunnerFunc func(t *testing.T)

// Run delegates down to the called on function `r`
func (r RunnerFunc) Run(t *testing.T) { r(t) }

// Case interface describes the set of things
// that can return a name string and are a Runner
type Case interface {
	Name() string
	Runner
}

// CommonCase implements Case
// Name returns the field CaseName
// Run delegates down to the embedded Runner
type CommonCase struct {
	name string
	Runner
}

// NewCase returns a pointer to a CommonCase, which
// implements the Case interface.
func NewCase(name string, runner Runner) *CommonCase {
	return &CommonCase{
		name:   name,
		Runner: runner,
	}
}

// Name returns the CaseName field string
func (c *CommonCase) Name() string { return c.name }

// Table is a slice of Case
type Table []Case

// Run takes a pointer to a testing.T and calls run on
// all the Case types in the Table slice.
func (table Table) Run(t *testing.T) {
	for _, testCase := range table {
		t.Run(testCase.Name(), testCase.Run)
	}
}
