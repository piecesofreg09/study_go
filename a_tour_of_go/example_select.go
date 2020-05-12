package main

import "fmt"
import "time"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			time.Sleep(1 * time.Nanosecond)
			fmt.Println(<-c)
		}
		//quit <- 1
	}()
	go func() {
		time.Sleep(10 * time.Nanosecond)
		quit <- 1
	}()
	fibonacci(c, quit)
	
}
