// Package fun implements is a demo to check documentation, and how it is
// working with gethub
package fun

// Init is the input number, to be used as struct insted of directly calling
// the functions for simplicity purposes
type Init struct {
	Number int
}

func main() {
	d := Init{Number: 3}
	d.Double()
	println(d.Number, "(which is 3 by 2)")
}

// Double function retuning the input value multiplied by 2
func (d *Init) Double() {
	d.Number = d.Number * 2
}

func ExampleDouble() {
	// Create new buffer for the image
	d := Init{Number: 3}
	d.Double()
	println(d.Number, "(which is 3 by 2)")
	// Output:
	// 6 (which is 3 by 2)
}
