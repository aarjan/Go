package fib

import "testing"

//import "fmt"

type fibTest struct {
	n        int
	expected int
}

var fibTests = []fibTest{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
	{8, 21},
}

func TestFib(t *testing.T) {
	//	t.Log("Running fib()")
	for _, tt := range fibTests {
		actual := Fib(tt.n)
		if actual != tt.expected {
			t.Errorf("Fib1(%d): Wanted %d, got %d", tt.n, tt.expected, actual)
		}
	}
}

func TestFib2(t *testing.T) {
	//	t.Log("Running fib2()")
	for _, tt := range fibTests {
		actual := Fib2(tt.n)
		if actual != tt.expected {
			t.Errorf("Fib2(%d): Wanted %d, got %d", tt.n, tt.expected, actual)
		}

	}
}

func TestFib3(t *testing.T) {
	for _, tt := range fibTests {
		actual := Fib3(tt.n)
		if actual != tt.expected {
			t.Errorf("Fib3(%d): Wanted %d, got %d", tt.n, tt.expected, actual)
		} else {
			t.Log(tt.n, tt.expected, actual)
		}
	}
}
func BenchmarkFib10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n <= b.N; n++ {
		Fib(10)
	}
}

func BenchmarkFib210(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		Fib2(10)
	}
}

func BenchmarkFib310(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		Fib3(10)
	}
}
