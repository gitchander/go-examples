package main

import (
	"math/rand"
	"strings"
)

var words = []IntVolume{
	{Data: 1, Volume: 194},
	{Data: 2, Volume: 1707},
	{Data: 3, Volume: 1899},
	{Data: 4, Volume: 1241},
	{Data: 5, Volume: 1059},
	{Data: 6, Volume: 909},
	{Data: 7, Volume: 888},
	{Data: 8, Volume: 668},
	{Data: 9, Volume: 522},
	{Data: 10, Volume: 418},
	{Data: 11, Volume: 222},
	{Data: 12, Volume: 140},
	{Data: 13, Volume: 133},
}

var wordsVolRand = NewVolRand(IntVolumes(words))

func RandWordLength(r *rand.Rand) int {
	index := wordsVolRand.RandIndex(r)
	return words[index].Data
}

// English Letter Frequency
// http://pi.math.cornell.edu/~mec/2003-2004/cryptography/subs/frequencies.html
var letters = []RuneVolume{
	{Data: 'e', Volume: 21912},
	{Data: 't', Volume: 16587},
	{Data: 'a', Volume: 14810},
	{Data: 'o', Volume: 14003},
	{Data: 'i', Volume: 13318},
	{Data: 'n', Volume: 12666},
	{Data: 's', Volume: 11450},
	{Data: 'r', Volume: 10977},
	{Data: 'h', Volume: 10795},
	{Data: 'd', Volume: 7874},
	{Data: 'l', Volume: 7253},
	{Data: 'u', Volume: 5246},
	{Data: 'c', Volume: 4943},
	{Data: 'm', Volume: 4761},
	{Data: 'f', Volume: 4200},
	{Data: 'y', Volume: 3853},
	{Data: 'w', Volume: 3819},
	{Data: 'g', Volume: 3693},
	{Data: 'p', Volume: 3316},
	{Data: 'b', Volume: 2715},
	{Data: 'v', Volume: 2019},
	{Data: 'k', Volume: 1257},
	{Data: 'x', Volume: 315},
	{Data: 'q', Volume: 205},
	{Data: 'j', Volume: 188},
	{Data: 'z', Volume: 128},
}

var lettersVolRand = NewVolRand(RuneVolumes(letters))

func RandLetter(r *rand.Rand) rune {
	index := lettersVolRand.RandIndex(r)
	return letters[index].Data
}

func RandWord(r *rand.Rand) string {
	n := RandWordLength(r)
	rs := make([]rune, n)
	for i := range rs {
		rs[i] = RandLetter(r)
	}
	return string(rs)
}

func RandLine(r *rand.Rand) string {
	n := RandInterval(r, 5, 15) // number of words
	var b strings.Builder
	for i := 0; i < n; i++ {
		if i > 0 {
			if x := r.Intn(10); x < 2 {
				b.WriteByte(',')
			}
			b.WriteByte(' ')
		}
		b.WriteString(RandWord(r))
	}
	b.WriteByte('.')
	return b.String()
}
