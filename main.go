package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	ch := make(chan int)
	a := []int{1, 2, 3}
	b := make([]int, 0, 4)

	for i := 1; i < len(a); i++ {
		wg.Add(1)
		go func(x, y int) {
			defer wg.Done()
			ch <- x + y
		}(a[i-1], a[i])

		res := <-ch
		b = append(b, res)
	}

	wg.Wait()
	fmt.Println(a)
	fmt.Println(b)
}
