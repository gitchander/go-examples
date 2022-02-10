package main

import (
	"fmt"
)

func generate(n int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			out <- i + 1
		}
		close(out)
	}()
	return out
}

func sqr(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	in := generate(20)
	out := sqr(in)
	for x := range out {
		fmt.Println(x)
	}
}
