package fun

import "testing"

// To be tested by running: go test -bench=.
func BenchmarkDouble(b *testing.B) {
	d := Init{Number: 3}
	// run the Double function b.N times
	for n := 0; n < b.N; n++ {
		d.Double()
	}
}
