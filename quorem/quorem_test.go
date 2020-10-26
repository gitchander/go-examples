package quorem

import (
	"math/rand"
	"testing"
	"time"
)

var seed int64

func init() {
	seed = time.Now().UTC().UnixNano()
}

func newRand() *rand.Rand {
	return rand.New(rand.NewSource(seed))
}

var (
	quo, rem   int
	quo1, rem1 int
	quo2, rem2 int
)

func TestQuoRem(t *testing.T) {
	const n = 10000
	for x := 0; x < n; x++ {
		for y := 1; y < n; y++ {

			quo1, rem1 = QuoRem1(x, y)
			quo2, rem2 = QuoRem2(x, y)

			if quo1 != quo2 {
				t.Fatalf("quoRem(%d, %d): quo1(%d) != quo2(%d)", x, y, quo1, quo2)
			}

			if rem1 != rem2 {
				t.Fatalf("quoRem(%d, %d): rem1(%d) != rem2(%d)", x, y, rem1, rem2)
			}

			if quo1*y+rem1 != x {
				t.Fatalf("quoRem(%d, %d): quo(%d), rem(%d)", x, y, quo1, rem1)
			}
		}
	}
}

// Local functions use local result variables (quo, rem)
func BenchmarkQuoRem1Local(b *testing.B) {
	r := newRand()
	var x, y int
	var quo, rem int
	for i := 0; i < b.N; i++ {
		x = r.Int()
		y = r.Int() + 1
		quo, rem = QuoRem1(x, y)
		_, _ = quo, rem
	}
}

func BenchmarkQuoRem2Local(b *testing.B) {
	r := newRand()
	var x, y int
	var quo, rem int
	for i := 0; i < b.N; i++ {
		x = r.Int()
		y = r.Int() + 1
		quo, rem = QuoRem2(x, y)
		_, _ = quo, rem
	}
}

func BenchmarkQuoRem1(b *testing.B) {
	r := newRand()
	var x, y int
	for i := 0; i < b.N; i++ {
		x = r.Int()
		y = r.Int() + 1
		quo, rem = QuoRem1(x, y)
	}
}

func BenchmarkQuoRem2(b *testing.B) {
	r := newRand()
	var x, y int
	for i := 0; i < b.N; i++ {
		x = r.Int()
		y = r.Int() + 1
		quo, rem = QuoRem2(x, y)
	}
}
