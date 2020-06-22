package unique

import (
	"math"
)

type Float32Slice []float32

func (p Float32Slice) Len() int          { return len(p) }
func (p Float32Slice) Hash(i int) uint64 { return uint64(math.Float32bits(p[i])) }
func (p Float32Slice) Swap(i, j int)     { p[i], p[j] = p[j], p[i] }

type Float64Slice []float64

func (p Float64Slice) Len() int          { return len(p) }
func (p Float64Slice) Hash(i int) uint64 { return math.Float64bits(p[i]) }
func (p Float64Slice) Swap(i, j int)     { p[i], p[j] = p[j], p[i] }

func Float32s(as []float32) []float32 {
	v := Float32Slice(as)
	n := Unique(v)
	return as[:n]
}

func Float64s(as []float64) []float64 {
	v := Float64Slice(as)
	n := Unique(v)
	return as[:n]
}
