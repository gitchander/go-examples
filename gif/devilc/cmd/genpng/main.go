package main

import (
	"log"

	"github.com/gitchander/go-examples/gif/devilc"
)

func main() {
	d := devilc.Devil{
		A: 90,
		B: 100,
	}
	if err := devilc.MakePNG("test.png", 512, 512, d); err != nil {
		log.Fatal(err)
	}
}
