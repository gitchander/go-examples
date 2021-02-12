package main

import "fmt"

type Swapper interface {
	Len() int
	Swap(i, j int)
}

type runeSlice []rune

func (p runeSlice) Len() int      { return len(p) }
func (p runeSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func Reverse(s Swapper) {
	i, j := 0, (s.Len() - 1)
	for i < j {
		s.Swap(i, j)
		i, j = i+1, j-1
	}
}

func main() {
	s := ".укнараб учрев ежкат есв я а ,укнаразопс олатсв ецнлоС"
	rs := []rune(s)
	Reverse(runeSlice(rs))
	fmt.Println(string(rs))
}
