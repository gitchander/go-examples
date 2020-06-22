package unique

func UniqueMapIndex(v Interface) map[uint64]int {
	m := make(map[uint64]int)
	n := v.Len()
	for i := 0; i < n; i++ {
		h := v.Hash(i)
		if _, ok := m[h]; !ok {
			m[h] = i
		}
	}
	return m
}

func UniqueMapIndexes(v Interface) map[uint64][]int {
	m := make(map[uint64][]int)
	n := v.Len()
	for i := 0; i < n; i++ {
		h := v.Hash(i)
		indexes := m[h]
		indexes = append(indexes, i)
		m[h] = indexes
	}
	return m
}
