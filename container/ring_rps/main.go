package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Stat struct {
	Wins  int
	Loses int
}

func (s Stat) String() string {
	return fmt.Sprintf("Wins - %d, Loses - %d", s.Wins, s.Loses)
}

func round(a Shape, stat *Stat) {
	b := nextShape()
	var r rune
	if a == b {
		r = '='
	} else {
		if a.Beats(b) {
			r = '>'
			stat.Wins++
		} else {
			r = '<'
			stat.Loses++
		}
	}
	fmt.Printf("%s %c %s\n", a, r, b)
}

type LineReader interface {
	ReadLine() (string, error)
}

type lineReader struct {
	br *bufio.Reader
}

func (p *lineReader) ReadLine() (string, error) {
	line, _, err := p.br.ReadLine()
	if err != nil {
		return "", err
	}
	return string(line), nil
}

func NewLineReader(r io.Reader) LineReader {
	return &lineReader{
		br: bufio.NewReader(r),
	}
}

func main() {

	var stat Stat

	lr := NewLineReader(os.Stdin)

	for {
		line, err := lr.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		command := strings.ToLower(line)

		switch command {
		case "q", "quit":
			return
		case "r", "rock":
			round(Rock, &stat)
		case "p", "paper":
			round(Paper, &stat)
		case "s", "scissors":
			round(Scissors, &stat)
		case "stat":
			fmt.Println(stat)
		default:
			fmt.Println("unknown command:", command)
		}
	}
}
