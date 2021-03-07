package main

import (
	"fmt"
	"math"
)

// Get initial parameters form user
func getInitialParameters() (float64, float64, float64) {
	var acceleration, initVelocity, initDisplacement float64

	fmt.Printf("Enter acceleration (m/s²): ")
	fmt.Scan(&acceleration)

	fmt.Printf("Enter initial velocity (m/s): ")
	fmt.Scan(&initVelocity)

	fmt.Printf("Enter initial displacement (m): ")
	fmt.Scan(&initDisplacement)

	return acceleration, initVelocity, initDisplacement
}

// Get time period from user
func getTime() float64 {
	var t float64
	fmt.Printf("Enter time (s): ")
	fmt.Scan(&t)
	return t
}

//Generate Displace function based on initial parameters
func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return a*math.Pow(t, 2)/2 + v0*t + s0
	}
}

func main() {
	displaceFn := GenDisplaceFn(getInitialParameters())
	t := getTime()

	fmt.Printf("Displacement after %.2f second(s): %f meter(s)\n", t, displaceFn(t))
}
