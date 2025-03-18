package main

import "testing"

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// run tests by running: `go test -v`
func TestMinBasic(t *testing.T) {
	result := Min(6, 9)
	if result != 6 {
		t.Errorf("Min(6, 9) = %d; want 6", result)
	}
}

func TestMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 0, 0},
		{4, 2, 2},
		{6, 9, 6},
		{-1, -10, -10},
		{-2, 5, -2},
	}

	for _, test := range tests {
		t.Run("find min between two integers", func(t *testing.T) {
			result := Min(test.a, test.b)
			if result != test.want {
				t.Errorf("Min(%d, %d)=%d; want %d", test.a, test.b, result, test.want)
			}
		})
	}

}

// run benchmarks by running: `go test -bench=.`
func BenchmarkMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Min(1, 2)
	}
}
