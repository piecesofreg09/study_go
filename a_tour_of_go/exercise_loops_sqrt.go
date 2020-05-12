package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	pre := z + 10
	for math.Abs(z - pre) > 1e-15 {
		pre = z
		fmt.Println(z)
		z -= (z*z - x) / (2*z) 
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
