Legs
====

Give your Golang table-driven tests some support (in roughly 50 LOC)

Requires go1.7+ (sub-test support required)

## About

A set of types / helper functions for writing table tests in Go.

All table-tests should be; is a test definition and a sequence of examples to run against that definition. This is a super small library to abstract away some of the common things that can be done. Like looping over the cases in the test and running them in a sub-test with a name.

## Examples

The example folder was adapted from Dave Cheneys [writing table driven tests in go](https://dave.cheney.net/2013/06/09/writing-table-driven-tests-in-go).

A table-test case is defined as something which implements the `legs.Case` interface. The Case interface provides a name for the case and a method to execute the test, provided with a pointer to a testing.T struct. 

From this `legs` can clear up some of the verbosity. Lets break the example down:

```go
type testCase struct {
	n        int // input
	expected int // expected result
}
```

Here we see all the important test specific fields required for a single row in our table-driven test.

```go
func (tt *testCase) Name() string {
	return fmt.Sprintf("Fib(%d) should equal %d", tt.n, tt.expected)
}

func (tt *testCase) Run(t *testing.T) {
	actual := Fib(tt.n)
	if actual != tt.expected {
		t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
	}
}
```

Now we have the implementations of Name and Run. Here we dynamically generate a name for the test case, based on the fields for the specific row. The run method actually performs a single run of the testCase itself. Performing the call to Fib and asserting the behavior using the provided testing.T pointer.

```go
func TestFib(t *testing.T) {
	legs.Table{
		newCase(1, 1),
		newCase(2, 1),
		newCase(3, 2),
		newCase(4, 3),
		newCase(5, 5),
		newCase(6, 8),
		newCase(7, 13),
	}.Run(t)
}
```

Finally we see where `legs` comes in to give our table some "structure" (excuse the punny name). A slice of `legs.Case` is a `legs.Table`, which happens to also implements `legs.Runner`. We can execute this table using the provided testing.T pointer on a call to Run. This generates the following output when run using `go test -v` in the example folder:

```
=== RUN   TestFib
=== RUN   TestFib/Fib(1)_should_equal_1
=== RUN   TestFib/Fib(2)_should_equal_1
=== RUN   TestFib/Fib(3)_should_equal_2
=== RUN   TestFib/Fib(4)_should_equal_3
=== RUN   TestFib/Fib(5)_should_equal_5
=== RUN   TestFib/Fib(6)_should_equal_8
=== RUN   TestFib/Fib(7)_should_equal_13
--- PASS: TestFib (0.00s)
    --- PASS: TestFib/Fib(1)_should_equal_1 (0.00s)
    --- PASS: TestFib/Fib(2)_should_equal_1 (0.00s)
    --- PASS: TestFib/Fib(3)_should_equal_2 (0.00s)
    --- PASS: TestFib/Fib(4)_should_equal_3 (0.00s)
    --- PASS: TestFib/Fib(5)_should_equal_5 (0.00s)
    --- PASS: TestFib/Fib(6)_should_equal_8 (0.00s)
    --- PASS: TestFib/Fib(7)_should_equal_13 (0.00s)
PASS
ok  	github.com/georgemac/legs/example	0.005s
```

## License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).
