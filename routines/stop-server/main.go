package main

import (
	"time"
)

func main() {

	s := OpenServer()
	time.Sleep(5 * time.Second)
	s.Close()
}
