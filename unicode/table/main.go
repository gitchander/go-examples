package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func GenerateUnicodeTable(fileName string, runeMin, runeMax rune) error {

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	bw := bufio.NewWriter(f)
	defer bw.Flush()

	i := 0
	for r := runeMin; r < runeMax; r++ {
		fmt.Fprintf(bw, "[%#U], ", r)

		if (i+1)%8 == 0 {
			bw.WriteByte('\n')
		}
		i++
	}

	return nil
}

func main() {
	err := GenerateUnicodeTable("unicode_table.txt", 64, 10074)
	if err != nil {
		log.Fatal(err)
	}
}

func printSomeRunes() {
	s := "\u1699 \u03A8 \u03CF \u090B \u0DA3 \u0FD5 \u2230"
	fmt.Println(s)
}
