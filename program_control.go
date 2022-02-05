package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
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

	// Switch statement for test value equality
	// Switch cases evaluate cases from top to bottom, stopping when a case succeeds. (Go only run the selected case, not all the cases that follow, implies adding break to the end of case body)
	fmt.Print("Go runs on ")
	// switch within a block-scoped variable declaration
	switch os := runtime.GOOS; os { // ; implies end of this statement, useful for writing multiple statement at a single line
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s. \n", os)
	}
	// fmt.Println(os) // undefined

	// Switch with condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// defer modifier keyword
	deferCall()
	deferStack()
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

// A defer statement defers the execution of a function until the surrounding function returns.
// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
func deferCall() {
	defer fmt.Println("world")
	// return
	fmt.Println("hello")
}

// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
func deferStack() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
