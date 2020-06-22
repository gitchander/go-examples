package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gitchander/go-examples/generics/instead/unique"
)

func main() {
	exampleInts()
	exampleStrings()
}

func exampleInts() {
	r := randNow()
	corpus := serialInts(10)
	v := randInts(r, corpus, 100)
	fmt.Println(v)
	a := unique.Ints(v)
	fmt.Println(a)
}

func exampleStrings() {
	r := randNow()
	corpus := []string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}
	v := randStrings(r, corpus, 100)
	fmt.Println(v)
	a := unique.Strings(v)
	fmt.Println(a)
}

func serialInts(n int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = i
	}
	return a
}

func randNow() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randInts(r *rand.Rand, corpus []int, n int) []int {
	vs := make([]int, n)
	for i := range vs {
		vs[i] = corpus[r.Intn(len(corpus))]
	}
	return vs
}

func randStrings(r *rand.Rand, corpus []string, n int) []string {
	vs := make([]string, n)
	for i := range vs {
		vs[i] = corpus[r.Intn(len(corpus))]
	}
	return vs
}
