package fibonacciseries

import (
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	testcases := []struct {
		name  string
		input int
		want  []int
	}{
		{name: "FibonacciSeriesL5", input: 5, want: []int{0, 1, 1, 2, 3}},
		{name: "FibonacciSeriesL1", input: 1, want: []int{0}},
		{name: "FibonacciSeriesL2", input: 2, want: []int{0, 1}},
		{name: "FibonacciSeriesL4", input: 4, want: []int{0, 1, 1, 2}},
		{name: "FibonacciSeriesL5", input: 0, want: []int{}},
	}

	for _, tc := range testcases {
		call := FibonacciSeries()
		got := call(tc.input)
		want := tc.want
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("Expected %q and got %q failed test case %s", want, got, tc.name)
		}
	}
}
