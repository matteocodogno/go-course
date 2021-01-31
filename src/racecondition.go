package main

import (
	"time"
)

/**
Race condition: Outcome depends on non-deterministic ordering.
Races occurs due to communication, it is very common there is some level of sharing between threads.

The following example shows 2 goroutines whose try to increase the same counter variable, the shared state which acts as communication layer, so the output is non-deterministic.

 */
var counter = 0

func increase() {
	counter ++
}

func main() {
	go increase()
	go increase()
	time.Sleep(time.Second)
}

