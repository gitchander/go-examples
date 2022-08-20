package main

import "fmt"

// type Number interface {
// 	int
// }

type Number interface {
	~int | ~int8
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

type Integer int8

func testMin() {
	var a, b Integer = 5, 3
	fmt.Println(Min(a, b))
}
