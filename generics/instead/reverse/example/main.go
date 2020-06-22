package main

import (
	"fmt"

	"github.com/gitchander/go-examples/generics/instead/reverse"
)

func main() {
	exampleInts()
}

func exampleInts() {
	a := serialInts(10)
	fmt.Println(a)
	reverse.Ints(a)
	fmt.Println(a)
}

func serialInts(n int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = i
	}
	return a
}
