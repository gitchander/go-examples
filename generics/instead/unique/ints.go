package unique

type IntSlice []int

func (p IntSlice) Len() int          { return len(p) }
func (p IntSlice) Hash(i int) uint64 { return uint64(p[i]) }
func (p IntSlice) Swap(i, j int)     { p[i], p[j] = p[j], p[i] }

func Ints(as []int) []int {
	v := IntSlice(as)
	n := Unique(v)
	return as[:n]
}
