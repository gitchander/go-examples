package unique

type Interface interface {
	Len() int
	Hash(i int) uint64
	Swap(i, j int)
}

func Unique(v Interface) int {
	var (
		m = make(map[uint64]struct{})
	)
	var (
		i = 0
		n = v.Len()
	)
	for i < n {
		h := v.Hash(i)
		if _, ok := m[h]; ok {
			n--
			v.Swap(i, n)
		} else {
			m[h] = struct{}{}
			i++
		}
	}
	return n
}
