package main

import (
	"fmt"
	"math"
)

func main() {
	a := -0.1
	x := math.Float64bits(a)
	fmt.Printf("%b\n", x)
}
