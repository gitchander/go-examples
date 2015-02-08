package divmod

import (
	"testing"
	"time"
)

func testDivMod(t *testing.T, divmod fDivMod, n uint64, i int) {

	begin := time.Now()

	for x := uint64(1); x < n; x++ {
		for y := uint64(1); y < x; y++ {
			divmod(x, y)
		}
	}

	t.Logf("v%d time -> %v\n", 1, time.Since(begin))
}

func TestDivModTime(t *testing.T) {

	const n = 20000
	testDivMod(t, v1_DivMod, n, 1)
	testDivMod(t, v2_DivMod, n, 2)
	testDivMod(t, v3_DivMod, n, 3)
}
