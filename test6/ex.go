package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v0, s0 float64) func (float64) float64{
	fn := func (t float64) float64{
		return 0.5 * a * math.Pow(t, 2) + v0 * t +s0
	}
	return fn
}

func main() {
	var a, v0, s0, t float64
	fmt.Print("Acceleration : ")
	_, _ = fmt.Scan(&a)
	fmt.Print("Initial Velocity : ")
	_, _ = fmt.Scan(&v0)
	fmt.Print("Initial Displacement : ")
	_, _ = fmt.Scan(&s0)
	fn := GenDisplaceFn(a, v0, s0)
	fmt.Print("Enter the value of time : ")
	_, _ = fmt.Scan(&t)
	fmt.Println(fn(t))
}