package main

// Unlike C, Go has no pointer arithmetic, meaning we can't add or subtract numbers on a pointer itself

import "fmt"

func main() {
	i, j := 42, 2701
	p := &i         // Store i's pointer to p
	fmt.Println(*p) // De-Referencing the pointer p and get the actual value
	*p = 21         // Set i's value
	fmt.Println(i)

	p = &j       // p now stores the memory address of variable j
	*p = *p / 37 // Go does not have pointer arithmetic
	fmt.Println(j)

	fmt.Println(Vertex{1, 2})

	myPoint := Vertex{1, 2}
	myPoint.X = 4
	myPointPtr := &myPoint
	myPointPtr.Y = 1e9 // implicit de-referencing myPointPtr, similar to myPointPtr->Y in C/C++
	fmt.Println(myPoint.X, myPointPtr.Y)

	myPointPtr = &Vertex{10, 20} // A pointer points to an instance of Vertex struct, with its X=10 and Y=20
	fmt.Println(*myPointPtr)
}

// Vertex is an 2-D tuple
type Vertex struct {
	X int
	Y int
}
