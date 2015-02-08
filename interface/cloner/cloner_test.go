package cloner

import (
	"fmt"
	"testing"
)

type Alpha struct {
	Id      int
	Message string
}

func (this *Alpha) Clone() Cloner {

	cloned := *this
	return &cloned
}

type Betha struct {
	Name     string
	Position int
}

func (this *Betha) Clone() Cloner {

	cloned := *this
	return &cloned
}

func TestCloner(t *testing.T) {

	a := Alpha{
		Id:      345,
		Message: "Message",
	}

	b := Betha{
		Name:     "asduqweruwisg",
		Position: 23534,
	}

	var cs = []Cloner{&a, &b}

	for _, c := range cs {
		clone := c.Clone()
		fmt.Printf("%+v\n", clone)
	}
}
