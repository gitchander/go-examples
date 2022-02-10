package main

import (
	"log"

	"github.com/gitchander/go-examples/gif/devilc"
)

func main() {
	err := devilc.MakeGif("devil-curve.gif")
	if err != nil {
		log.Fatal(err)
	}
}
