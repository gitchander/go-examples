package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	data := make([]byte, binary.MaxVarintLen64)
	r := newRand()
	for i := 0; i < 1000; i++ {
		a := randInt64(r)
		n := binary.PutVarint(data, a)
		fmt.Printf("%d: [%x]\n", a, data[:n])
		b, _ := binary.Varint(data)
		if a != b {
			log.Fatalf("%d != %b", a, b)
		}
	}
}

func newRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randBool(r *rand.Rand) bool {
	return (r.Int() & 1) == 1
}

func randInt64(r *rand.Rand) int64 {
	a := r.Int63() >> r.Intn(63)
	if randBool(r) { // random sign
		a = -a
	}
	return a
}
