package main

import (
	"fmt"
	"time"

	"github.com/gitchander/go-examples/lychrel"
)

func main() {

	start := time.Now()

	const base = 10
	table := lychrel.NewTable(base)

	//n := NewNumber(table, 10)
	n := lychrel.NewNumber(table, 196)
	//n := NewNumber(table, 57)
	//n := NewNumber(table, 1186060307891929990)

	var iter, number int
	const count = 1000

	for i := 0; i < 10; i++ {

		number = lychrel.LychrelTest(n, count)
		iter += number
		if number < count {
			break
		}

		fmt.Printf("Iterations: %d; count digits: %d\n", iter, n.CountDigits())
	}

	fmt.Printf("time duration: %v\n", time.Since(start))

	if number < count {
		fmt.Printf("Is Lychrel number after %d iterations\n", iter)

		if n.CountDigits() < 200 {
			fmt.Printf("%s\n", n)
		}
	}

	rd := lychrel.FindLychrelNumbers(10000, 10)
	fmt.Println(rd)

	//bs := []byte{23, 1, 2, 53, 4, 1, 23}
	//fmt.Printf("%v is palindrome: %v\n", bs, IsPalindrome(bs))

}
