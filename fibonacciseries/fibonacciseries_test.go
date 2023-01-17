package fibonacciseries

import (
	"reflect"
	"testing"
)

func TestFibonacciNumberL5(t *testing.T) {
	call := FibonacciSeries() //Slice
	got := call(5)
	want := []int{0, 1, 1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %q and got %q", want, got)
	}
}

func TestFibonacciNumberL2(t *testing.T) {
	call := FibonacciSeries() //Slice
	got := call(2)
	want := []int{0, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %q and got %q", want, got)
	}
}

func TestFibonacciNumberL1(t *testing.T) {
	call := FibonacciSeries() //Slice
	got := call(1)
	want := []int{0}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %q and got %q", want, got)
	}
}
