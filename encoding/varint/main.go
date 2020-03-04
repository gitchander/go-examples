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
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 1000; i++ {
		a := r.Int63() >> r.Intn(63)
		if (r.Int() & 1) == 1 {
			a = -a
		}
		n := binary.PutVarint(data, a)
		fmt.Printf("%d: [%x]\n", a, data[:n])
		b, _ := binary.Varint(data)
		if a != b {
			log.Fatalf("%d != %b", a, b)
		}
	}
}
