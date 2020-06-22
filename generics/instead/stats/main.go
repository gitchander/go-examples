package main

import (
	"fmt"
)

func main() {
	a := []int{4, 36, 45, 50, 75}
	st := CalcStatistics(IntSlice(a))
	fmt.Printf("%+v\n", st)
}

type Interface interface {
	Len() int
	Value(i int) float64
}

type Statistics struct {
	Mean float64
	MSE  float64
}

func CalcStatistics(v Interface) Statistics {

	n := v.Len()

	var (
		sum    float64
		sumSqr float64
	)

	for i := 0; i < n; i++ {

		x := v.Value(i)

		sum += x
		sumSqr += x * x
	}

	var (
		mean = sum / float64(n)
		mse  = sumSqr / float64(n)
	)

	return Statistics{
		Mean: mean,
		MSE:  mse,
	}
}

type IntSlice []int

func (p IntSlice) Len() int            { return len(p) }
func (p IntSlice) Value(i int) float64 { return float64(p[i]) }
