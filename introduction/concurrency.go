package main

import (
	"fmt"
	"time"
)

func main() {
	// Goroutines: go f(x, y, z) starts a new goroutine running f(x, y, z)
	go say("world") // Execute this function on a different thread and immediately move on executing functions afterwards
	say("Hello")

	s := []int{7, 2, 8, -9, 4, 0}

	// Channels (similar to a Pipe in UNIX software/script programming)
	fmt.Println()
	c := make(chan int)     // make a channel
	go sum(s[:len(s)/2], c) // [7,2,8]  c <- 17
	go sum(s[len(s)/2:], c) // [-9,4,0] c <- -5
	x, y := <-c, <-c        // receive from c (The data flows in the direction of the arrow.)
	fmt.Println(x, y, x+y)  // Once both goroutines have completed their computation, it calculates the final result.
	// By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.
	// This is similar to await modifier keyword, such as a fetchAPI async function will not continue to parse the response data until the request (executed asynchronously) is completed
	// x1, y1 := <-c, <-c // If I want to receive channels again, it will throw an error fatal error: all goroutines are asleep - deadlock!
	// fmt.Println(x1, y1, x1+y1)

}

// JavaScript's setTimeOut function to delay task execution, and JavaScript has coroutine concept (await modifier keyword in an async function)
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Println(sum)
	c <- sum // send sum to c
}
