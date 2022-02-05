package main

import (
	"fmt"
	"math"
)

// Go only has for-loop, no other program control flows
func main() {
	// For-loop
	sum := 0
	sum2 := 1
	for i := 0; i < 10; i++ {
		sum += i
	}

	// The init and post statements are optional (equivalent to while loop)
	for sum2 < 1000 {
		sum2 += sum2
	}

	// Omitting any condition will cause for-loop to run forever
	for {
		break
	}

	fmt.Println(sum)
	fmt.Println(sum2)

	fmt.Println(sqrt(2), sqrt(4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

func sqrt(x float64) string {
	// if statement
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// Block-scoped variable inside if
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// fmt.Println(v) // undefined
	return lim
}
