package main

import (
	"fmt"
	"math"
)

func isChanging(val float64) bool {
	if math.Abs(val) < 0.000000001 {
		return false
	} else {
		return true
	}
}

func Sqrt(x float64) float64 {
	z := 1.0
	diff := math.MaxFloat64
	for count := 1; isChanging(diff); count++ {
		fmt.Println("", count, z, diff)
		diff = (z*z - x) / (2 * z)
		fmt.Println("after cal= ", diff)
		z -= diff
	}
	return z
}

func main() {
	fmt.Println(Sqrt(5555))
	fmt.Println("actual sqrt = ", math.Sqrt(5555))
}
