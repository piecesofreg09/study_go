package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkR(t, ch)
	close(ch)
}

func WalkR(t *tree.Tree, ch chan int) {
	if t != nil {
		WalkR(t.Left, ch)
		ch <- t.Value
		WalkR(t.Right, ch)
	}
	
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	for i := 0; i < 10; i++ {
		x1 := <- ch1
		x2 := <- ch2
		if x1 != x2 {
			return false
		}
	}
	
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	fmt.Println(tree.New(1))
	
	t1 := Same(tree.New(1), tree.New(1))
	fmt.Println(t1)
	t2 := Same(tree.New(1), tree.New(2))
	fmt.Println(t2)
	t3 := Same(tree.New(2), tree.New(2))
	fmt.Println(t3)
}
