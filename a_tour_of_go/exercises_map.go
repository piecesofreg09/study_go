package main

import (
	"golang.org/x/tour/wc"
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	var x map[string]int
	x = make(map[string]int)
	sarray := strings.Fields(s)
	for _, v := range sarray {
		fmt.Println(v)
		_, ok := x[v]
		if ok {
			x[v] += 1
		} else {
			x[v] = 1
		}
	}
	return x
}

func main() {
	wc.Test(WordCount)
}
