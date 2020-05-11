package main

import "golang.org/x/tour/pic"
import "math"

func Pic(dx, dy int) [][]uint8 {
	
	var data = make([][]uint8, dy)
	
	for i := range data {
		var temp = make([]uint8, dx)
		for j := range temp {
			temp[j] = uint8(i * j * int(math.Sqrt(float64(j))))
		}
		data[i] = temp
	}
	
	return data
}

func main() {
	pic.Show(Pic)
}
