package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	stop_margin := 0.0000000000001

	z := 1.0
	prev_z := 0.0

	for i := 0; i < 100; i++ {
		prev_z = z
		z -= (z*z - x) / (2 * z)

		if math.Abs(prev_z-z) < stop_margin {
			fmt.Println("Iteration: ", i, "| change is less than: ", stop_margin, "| Value: ", z)
			break
		}

		fmt.Println("Iteration: ", i, "| Value: ", z)
	}
	fmt.Println("Actual value: ", math.Sqrt(x))
	fmt.Println("Final value:  ", z)
	return z
}

func main___() {
	Sqrt(2)
}
