package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(20 * time.Millisecond)
	//boom := time.After(100 * time.Millisecond)
	count := 0
	for {
		select {
		case <-tick:
			fmt.Println(time.Now().UnixNano())
			fmt.Println("tick.")
			count += 1
			if count == 4 {
				return
			}
			time.Sleep(50 * time.Nanosecond)
		//case <-boom:
		//	fmt.Println("BOOM!")
		//	return
		default:
			fmt.Println(time.Now().UnixNano())
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}


/* Output will be:
1257894000000000000
    .
1257894000050000000
tick.
1257894000050000050
    .
1257894000100000050
tick.
1257894000100000100
    .
1257894000150000100
tick.
1257894000150000150
    .
1257894000200000150
tick.

so I guess each tick value that the channel created is not buffed.
In the above time only 4 ticks are counted but there should be 20 
that emitted.

*/
