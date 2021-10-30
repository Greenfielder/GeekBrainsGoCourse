package fibonachi

import "testing"

type testCases struct {
	input  uint
	result uint
}

var cases = []testCases{
	{input: 0, result: 0},
	{input: 1, result: 1},
	{input: 2, result: 1},
	{input: 3, result: 2},
	{input: 5, result: 5},
	{input: 10, result: 55},
	{input: 15, result: 610},
	{input: 20, result: 6765},
	{input: 30, result: 832040},
	{input: 40, result: 102334155},
}

func TestFibonachi(t *testing.T) {
	for _, c := range cases {
		if res := FibonachiCached(c.input); res != c.result {
			t.Fatalf("Input: %d, Correct value is %d, your value is %d", c.input, c.result, res)
		}

	}
}
