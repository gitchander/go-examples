package main

import (
	"fmt"
)

func main() {
	testRandLines()
}

func testRandLines() {
	r := NewRandNow()
	for i := 0; i < 10; i++ {
		fmt.Println(RandLine(r))
	}
}

// RPS - Rock Paper Scissors
func testRPS() {
	var vs = []StringVolume{
		{Data: "rock", Volume: 1},
		{Data: "paper", Volume: 1},
		{Data: "scissors", Volume: 1},
	}
	vr := NewVolRand(StringVolumes(vs))
	r := NewRandNow()
	for i := 0; i < 100; i++ {
		index := vr.RandIndex(r)
		fmt.Println(vs[index].Data)
	}
}
