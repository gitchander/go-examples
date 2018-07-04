package main

import "fmt"

// https://en.wikipedia.org/wiki/Palindrome

func main() {
	samplesLatin := []string{
		"saippuakivikauppias",
		"A man, a plan, a canal, Panama!",
		"Was it a car or a cat I saw?",
		"Able was I ere I saw Elba",
		"Madam, I'm Adam",
		"Never odd or even",
		"Doc, note: I dissent. A fast never prevents a fatness. I diet on cod",
		"T. Eliot, top bard, notes putrid tang emanating, is sad; I'd assign it a name: gnat dirt upset on drab pot toilet.",
	}
	printSamples(samplesLatin, prepareLatin)

	samplesCirillic := []string{
		"А роза упала на лапу Азора",
		"Аргентина манит негра",
		"Я иду с мечем судия",
		"Коса налетела на сак.",
	}
	printSamples(samplesCirillic, prepareCirillic)
}

func printSamples(samples []string, prepare func(string) []rune) {
	for _, sample := range samples {
		rs := prepare(sample)
		fmt.Printf("[%s]", sample)
		if IsPalindrome(RuneSlice(rs)) {
			fmt.Println(" - is palindrome!")
		} else {
			fmt.Println(" - isn't palindrome")
		}
	}
}

func prepareLatin(s string) []rune {
	var rs []rune
	for _, r := range s {
		if ('a' <= r) && (r <= 'z') {
			rs = append(rs, r)
		}
		if ('A' <= r) && (r <= 'Z') {
			r = (r - 'A') + 'a'
			rs = append(rs, r)
		}
	}
	return rs
}

func prepareCirillic(s string) []rune {
	var rs []rune
	for _, r := range s {
		if ('а' <= r) && (r <= 'я') {
			rs = append(rs, r)
		}
		if ('А' <= r) && (r <= 'Я') {
			r = (r - 'А') + 'а'
			rs = append(rs, r)
		}
	}
	return rs
}

type RuneSlice []rune

func (p RuneSlice) Len() int             { return len(p) }
func (p RuneSlice) Equals(i, j int) bool { return p[i] == p[j] }

type Interface interface {
	Len() int

	// Match
	Equals(i, j int) bool
}

func IsPalindrome(v Interface) bool {
	var i, j = 0, v.Len() - 1
	for i < j {
		if !v.Equals(i, j) {
			return false
		}
		i, j = i+1, j-1
	}
	return true
}
