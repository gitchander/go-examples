package main

import (
	"fmt"
	"time"

	"github.com/gitchander/go-examples/lychrel"
)

func numberTest(v uint64) {

	start := time.Now()

	fmt.Printf("test number %v:\n", v)

	n := lychrel.NewNumber(v)

	var iter, number int
	const count = 1000

	for i := 0; i < 10; i++ {

		number = lychrel.LychrelTest(n, count)
		iter += number
		if number < count {
			break
		}

		fmt.Printf("\titerations: %d; count digits: %d\n", iter, n.CountDigits())
	}

	if number < count {
		fmt.Printf("\tlychrel number after iterations: %d\n", iter)

		if n.CountDigits() < 200 {
			fmt.Printf("\tpalindrome: %s\n", n)
		}
	}

	fmt.Printf("\ttime duration: %v\n", time.Since(start))

	fmt.Println()
}

func main() {

	numberTest(10)
	numberTest(57)
	numberTest(196)
	numberTest(1186060307891929990)

	rd := lychrel.FindLychrelNumbers(10000, 10)
	fmt.Println(rd)
}
