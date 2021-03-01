package main

import (
	"math/rand"
)

//------------------------------------------------------------------------------
// Percentage:
// https://en.wikipedia.org/wiki/Percentage

// Volume fraction:
// https://en.wikipedia.org/wiki/Volume_fraction#Volume_percent
//------------------------------------------------------------------------------

// Volumer, Interface
type Volumes interface {
	Len() int
	Volume(i int) int
}

// RandVol, VolRand, VolumeRand
type VolRand struct {
	parts []int
	sum   int
}

func NewVolRand(vs Volumes) *VolRand {

	n := vs.Len()
	parts := make([]int, n)
	var sum int
	for i := 0; i < n; i++ {
		volume := vs.Volume(i)
		if volume > 0 {
			sum += volume
		}
		parts[i] = sum
	}

	return &VolRand{
		parts: parts,
		sum:   sum,
	}
}

func (p *VolRand) RandIndex(r *rand.Rand) int {
	x := r.Intn(p.sum)
	for i, part := range p.parts {
		if x < part {
			return i
		}
	}
	return -1
}

type IntVolume struct {
	Data   int
	Volume int
}

type RuneVolume struct {
	Data   rune
	Volume int
}

type StringVolume struct {
	Data   string
	Volume int
}

type IntVolumes []IntVolume

func (v IntVolumes) Len() int         { return len(v) }
func (v IntVolumes) Volume(i int) int { return v[i].Volume }

type RuneVolumes []RuneVolume

func (v RuneVolumes) Len() int         { return len(v) }
func (v RuneVolumes) Volume(i int) int { return v[i].Volume }

type StringVolumes []StringVolume

func (v StringVolumes) Len() int         { return len(v) }
func (v StringVolumes) Volume(i int) int { return v[i].Volume }
