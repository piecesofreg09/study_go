package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64
func (e ErrNegativeSqrt) Error() string {
	return "cannot Sqrt negative number:" + fmt.Sprint(float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		y := ErrNegativeSqrt(x)
		return x, y
	} else {
		z := x / 2
		pre := z + 10
		for math.Abs(z - pre) > 1e-15 {
			pre = z
			fmt.Println(z)
			z -= (z*z - x) / (2*z) 
		}
		return z, nil
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
