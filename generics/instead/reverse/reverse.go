package reverse

type Swapper interface {
	Len() int
	Swap(i, j int)
}

func Reverse(s Swapper) {
	var (
		i = 0
		j = s.Len() - 1
	)
	for i < j {
		s.Swap(i, j)
		i++
		j--
	}
}

type IntSlice []int

func (p IntSlice) Len() int      { return len(p) }
func (p IntSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type StringSlice []string

func (p StringSlice) Len() int      { return len(p) }
func (p StringSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func Ints(a []int) {
	Reverse(IntSlice(a))
}

func Strings(a []string) {
	Reverse(StringSlice(a))
}
