package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Hello, Round!")

	fmt.Println(math.Round(-0.1)) // -0
	fmt.Println(math.Round(0.1))  // 0

	fmt.Println(math.Round(0.4999999999999999))  // 0
	fmt.Println(math.Round(0.49999999999999999)) // 1

	fmt.Println(math.Round(-0.4999999999999999))  // -0
	fmt.Println(math.Round(-0.49999999999999999)) // -1

	fmt.Println(math.Round(-1000.4999999999999))  // -1000
	fmt.Println(math.Round(-1000.49999999999999)) // -1001

	fmt.Println(math.Round(math.Inf(1)))  // +Inf
	fmt.Println(math.Round(math.Inf(-1))) // -Inf

	fmt.Println(math.Round(math.NaN())) // NaN

	fmt.Println(math.Round(Zero(false))) // 0
	fmt.Println(math.Round(Zero(true)))  // -0
}

func Zero(negative bool) float64 {
	zero := 0.0
	if negative {
		zero = -zero
	}
	return zero
}

func examplesNegativeZero() {

	// var.1
	fmt.Println(1 / math.Inf(-1))

	// var.2
	zero := 0.0
	fmt.Println(-zero)

	// var.3
	fmt.Println(math.Copysign(0, -1))
}
