package divmod

type fDivMod func(dividend, divisor uint64) (quotient, remainder uint64)

func v1_DivMod(dividend, divisor uint64) (quotient, remainder uint64) {

	quotient = dividend / divisor
	remainder = dividend % divisor

	return
}

func v2_DivMod(dividend, divisor uint64) (quotient, remainder uint64) {

	quotient = dividend / divisor
	remainder = dividend - quotient*divisor

	return
}

func v3_DivMod(x, y uint64) (q, r uint64) {
	sh := 0
	for (y+y > y) && (y+y <= x) {
		sh++
		y <<= 1
	}
	for ; sh >= 0; sh-- {
		q <<= 1
		if x >= y {
			x -= y
			q |= 1
		}
		y >>= 1
	}
	return q, x
}
