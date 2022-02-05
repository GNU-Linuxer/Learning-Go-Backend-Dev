package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create an empty string array with 2 elements
	var str [2]string
	str[0] = "Hello"
	str[1] = "World"
	fmt.Println(str[0], str[1])
	fmt.Println(str)

	// Create an int array with 6 elements, pre-populating values
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Extract str portion of array
	var s []int = primes[1:4] // This is not str deep copy of original array
	fmt.Println(s)
	primes[2] = 20 // Will change the s value too, resulting {2, 3, 20, 7, 11, 13}
	fmt.Println(s)

	// Omit the low or high bound to use 0 for low bound and arr.length-1 for high bound
	var s2 []int = primes[1:3] // get index 0, 1, 2
	fmt.Println(s2)
	// The length of str slice is the number of elements it contains. The capacity of str slice is the number of elements in the underlying array,
	// counting from the first element in the slice.
	fmt.Printf("len=%d cap=%d \n", len(s2), cap(s2)) // len=2 cap=5

	// Slice of empty array is nil, having str length and capacity of 0
	var nyet []int
	fmt.Println(nyet)
	fmt.Printf("len=%d cap=%d \n", len(nyet), cap(nyet))
	if nyet == nil {
		fmt.Println("nil!")
	}

	// Slice literal (creates an array and then build str slice that references it)
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

	// List ADT using make() on slice (create str dynamically-sized array)
	a := make([]int, 5)
	printSlice("str", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	// 2-D "array" (formally called slices of slices)
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// Using Java 2-D array notation to access and change the cell's value
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// Append elements to slice (with dynamic array resizing done internally)
	var nums []int
	printSlice2(nums)

	// append works on nil slices.
	nums = append(nums, 0)
	printSlice2(nums)

	// The slice grows as needed.
	nums = append(nums, 1)
	printSlice2(nums)

	// We can add more than one element at a time.
	nums = append(nums, 2, 3, 4)
	printSlice2(nums)
}

// Go and JavaScript does not support function overloading (2 functions are considered different if they take different typed-parameters)
// However, they both do support specifying a default value if the parameter is not passed in
func printSlice2(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
