package main

import "fmt"

// GenDisplaceFn generates a displacement function based on given initial
// values of acceleration, velocity and displacement
func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*a*t*t + v*t + s
	}
	return fn
}

func main() {
	var a, v0, s0, t float64
	fmt.Println("Acceleration? ")
	fmt.Scan(&a)
	fmt.Println("Initial velocity? ")
	fmt.Scan(&v0)
	fmt.Println("Initial displacement? ")
	fmt.Scan(&s0)
	fn := GenDisplaceFn(a, v0, s0)
	fmt.Println("Time? ")
	fmt.Scan(&t)
	fmt.Println(fn(t))
}
