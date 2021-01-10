package main

import (
	"fmt"
	"math"
)

func main() {
	var a, v0, s0, t float64

	fmt.Printf("Enter acceleration: \n")
	fmt.Scan(&a)
	fmt.Printf("Enter initial velocity: \n")
	fmt.Scan(&v0)
	fmt.Printf("Enter initial displacement: \n")
	fmt.Scan(&s0)
	fmt.Printf("Enter time: \n")
	fmt.Scan(&t)

	fn := GenDisplaceFn(a, v0, s0)

	fmt.Printf("Displacement %f", fn(t))
}

func GenDisplaceFn(a, v0, s0 float64) func (float64) float64 {
	return func(t float64) float64 {
		return (0.5 * a * math.Pow(t, 2)) + v0 * t + s0
	}
}