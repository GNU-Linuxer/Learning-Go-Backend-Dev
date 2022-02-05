package main

import "fmt"

func main() {
	i, j := 42, 2701
	p := &i         // Store i's pointer to p
	fmt.Println(*p) // De-Referencing the pointer p and get the actual value
	*p = 21         // Set i's value
	fmt.Println(i)

	p = &j // p now stores the memory address of variable j
	*p = *p / 37
	fmt.Println(j)

}
