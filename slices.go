package main

import "fmt"

func main() {

	// Create an empty string array with 2 elements
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// Create an int array with 6 elements, pre-populating values
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Extract a portion of array
	var s []int = primes[1:4] // This is not a deep copy of original array
	fmt.Println(s)
	primes[2] = 20 // Will change the s value too, resulting {2, 3, 20, 7, 11, 13}
	fmt.Println(s)

	// Omit the low or high bound to use 0 for low bound and arr.length-1 for high bound
	var s2 []int = primes[1:3] // get index 0, 1, 2
	fmt.Println(s2)
	// The length of a slice is the number of elements it contains. The capacity of a slice is the number of elements in the underlying array,
	// counting from the first element in the slice.
	fmt.Printf("len=%d cap=%d \n", len(s2), cap(s2)) // len=2 cap=5

	// Slice of empty array is nil, having a length and capacity of 0
	var nyet []int
	fmt.Println(nyet)
	fmt.Printf("len=%d cap=%d \n", len(nyet), cap(nyet))
	if nyet == nil {
		fmt.Println("nil!")
	}

	// Slice literal (creates an array and then build a slice that references it)
	q := []int{2, 3, 5, 7, 11, 13} // Note that there's nothing in []
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	t := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(t)

	// List ADT using slice (create a dynamically-sized array)
}
