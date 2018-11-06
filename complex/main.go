package main

import (
	"fmt"
	"math/cmplx"
)

func main() {

	fmt.Printf("%T\n", complex(float32(1), float32(2)))

	var c1 complex64 = complex(0.5, 1)
	var c2 complex128 = complex(3.1, -7)

	fmt.Printf("%v = %.2f + i * (%.2f)\n", c1, real(c1), imag(c1))
	fmt.Printf("%v = %.2f + i * (%.2f)\n", c2, real(c2), imag(c2))

	// convert
	c2 = complex128(c1)
	c1 = complex64(c2)

	z := complex128(complex(-1, 0))
	r, θ := cmplx.Polar(z)
	fmt.Println(r, θ)

	r, θ = cmplx.Polar(complex(1, 2))
	fmt.Println(r, θ)
}
