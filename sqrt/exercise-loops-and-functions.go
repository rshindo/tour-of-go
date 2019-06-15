package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	var z = x
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println("math.Sqrt => ", math.Sqrt(2))
}
