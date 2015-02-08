package main

import (
	"fmt"
	"math/big"
)

func Fibb() func() uint64 {
	var a, b uint64 = 0, 1
	return func() uint64 {
		a, b = b, a+b
		return a
	}
}

func BigFibb() func() *big.Int {

	a := new(big.Int)
	b := new(big.Int)
	c := new(big.Int)

	a.SetInt64(0)
	b.SetInt64(1)

	return func() *big.Int {
		c.Add(a, b)
		a.Set(b)
		b.Set(c)
		return a
	}
}

func main() {

	f := BigFibb()
	for i := 0; i < 1000; i++ {
		fmt.Printf("%s\n", f())
	}
}
