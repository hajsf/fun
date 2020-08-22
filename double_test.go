package fun

import "testing"

// To be tested by running: go test -v
func TestDouble(t *testing.T) {
	d := Init{Number: 3}
	d.Double()
	want := 6
	if got := d.Number; got != want {
		println("test failed")
	} else {
		println("test passed")
	}
}
