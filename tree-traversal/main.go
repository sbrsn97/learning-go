package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch := make(chan int)
	ch2 := make(chan int)

	go func() {
		Walk(t1, ch)
		close(ch)
	}()

	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	for {
		v1, ok1 := <-ch
		v2, ok2 := <-ch2

		if !ok1 && !ok2 {
			break
		}

		if ok1 != ok2 || v1 != v2 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("1 ve 2 aynı mı?:", Same(tree.New(1), tree.New(2)))

	fmt.Println("1 ve 1 aynı mı?:", Same(tree.New(1), tree.New(1)))
}
