package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fi := 1
	pre := 0
	
	return func() int {
		tempFi := fi
		fi += pre
		tempPre := pre
		pre = tempFi
		return tempPre
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
