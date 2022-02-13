package main

import (
	"fmt"
	"math"
	"strconv"
)

// A type implements an interface by implementing its methods.
type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() { // It does not contain explicit implementation
	fmt.Println(t.S)
}

// Type switch
func do(i interface{}) {
	switch v := i.(type) {
	case int: // a variable v holds the value i with type of int
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string: // a variable v holds the value i with type of string
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default: // the variable v is of the same interface type and value as i
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	v := Vertex2{3, 4}
	fmt.Println(v.Abs())

	// Implies (&v).Scale(10)
	v.Scale(10) // Pass by reference not by value; hence the original v would change (program behavior will change if we remove * from *Vertex2)
	fmt.Println(v.Abs())

	var i I = T{"hello"} // Comparable to List<String> i = new ArrayList<String>();
	i.M()

	do(21)
	do("hello")
	do(true)

	// Fetch error (strconv.Atoi will return the value and error), and we encounter an error if it's not nil
	str := "10"
	if s, err := strconv.Atoi(str); err == nil {
		fmt.Printf("%T, %v", s, s)
	}
}

type Vertex2 struct {
	X, Y float64
}

// A method is just a function with a receiver argument.
func (v Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// This is a more typical way to write the above Abs() function:
func _Abs(v Vertex2) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Pointer receiver function
func (v *Vertex2) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
