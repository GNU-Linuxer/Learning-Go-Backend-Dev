package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

// Can also use
// import "fmt"
// import "math"

// Global variables (package-level)
// Initial values when not explicitly stated: 0 for numeric types, false for boolean, and "" (the empty string) for string
var c, python, java = true, false, "no!" // Inferred type from initialization value

// Basic Type
// Default type when explicitly initialized: int for whole number, float64 for decimal, complex128 for complex number
// Note Go can sometimes catch Integer overflow and thus throw an error instead of silently proceeds
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1            // Bitwise operation on this unsigned 64-bit integer
	z      complex128 = cmplx.Sqrt(-5 + 12i) // Store a complex number
)

const Pi = 3.14

// Each time you run the example program rand.Intn will return the same number, to see a different number, seed the number generator; see rand.Seed.
func main() {
	// Functions and imports
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi) // In Go, a name is exported if it begins with a capital letter.
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))

	// Local variable (function-level)
	var i, j int = 1, 2
	k := 3 // declare and initialize
	fmt.Println(i, j, k, c, python, java)

	fmt.Printf("type: %T Value: %v \n", ToBe, ToBe)
	fmt.Printf("type: %T Value: %v \n", MaxInt, MaxInt)
	fmt.Printf("type: %T Value: %v \n", z, z)

	// Type conversion
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

func add(x int, y int) int { // Kotlin-style
	return x + y
}

func swap(x, y string) (string, string) { // implies (x string, y string)
	return y, x
}

// Named return values
func split(sum int) (x, y int) { // similar to other language's: type result = ****; do something; return result;
	x = sum * 4 / 9
	y = sum - x
	// A return statement without arguments returns the named return values. This is known as a "naked" return.
	return // Equivalent to: return x, y
}
