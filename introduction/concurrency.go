package main

import (
	"fmt"
	"sync"
	"time"
)

// Binary Tree Node
type TreeNode struct {
	Left  *TreeNode // Store a pointer to the left tree
	Value int
	Right *TreeNode // Store a pointer to the right tree
}

// SafeCounter is safe to use concurrently, using Mutex
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

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

	// Buffer with capacity (called channels)
	ch := make(chan int, 2) // Make a buffer with a capacity of 2
	ch <- 1
	ch <- 2
	// ch <- 3 // Throws an exception: all goroutines are asleep - deadlock!
	fmt.Println(<-ch) // Queue-like behaviors: FIFO order when removing elements from buffer, removing the earliest element entering the buffer
	fmt.Println(<-ch)

	fmt.Println()

	d := make(chan int, 5)
	go fibonacci(cap(d), d) // cap(d) is the channel capacity of d;
	for i := range d {      // Sender will tell the receiver that no more values are coming
		fmt.Println(i)
	}

	safeCounter := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go safeCounter.Increment("somekey")
	}

	time.Sleep(2 * time.Second) // If we comment it out, the value can be random because we're still calling the Increment function (but program won't crash because it will wait for one Increment to unlock the resource
	fmt.Println(safeCounter.Value("somekey"))

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

// Sender can choose to close a channel, so sender can tell the receiver that there are no more values coming to this channel
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// Go's standard library provides mutual exclusion with sync.Mutex and its two methods:
//	Lock
//	Unlock

// Increment increments the counter for the given key.
func (c *SafeCounter) Increment(key string) {
	c.mu.Lock() // Program will pause here until the resource is unlocked
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock() // Use defer to ensure the mutex will be unlocked
	return c.v[key]
}
