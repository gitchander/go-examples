package unique

import (
	"hash"
	"hash/fnv"
)

func Strings(as []string) []string {
	v := StringSlice(as)
	n := Unique(v)
	return as[:n]
}

func StringSlice(vs []string) Interface {
	return &stringsUniquer{
		h:  fnv.New64a(),
		vs: vs,
	}
}

type stringsUniquer struct {
	h  hash.Hash64
	vs []string
}

func (p *stringsUniquer) Len() int {
	return len(p.vs)
}

func (p *stringsUniquer) Hash(i int) uint64 {
	p.h.Reset()
	p.h.Write([]byte(p.vs[i]))
	return p.h.Sum64()
}
func (p *stringsUniquer) Swap(i, j int) {
	p.vs[i], p.vs[j] = p.vs[j], p.vs[i]
}
